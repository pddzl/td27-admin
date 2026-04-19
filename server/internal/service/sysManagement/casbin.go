package sysManagement

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"go.uber.org/zap"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
	casbinpkg "server/internal/pkg/casbin"
)

type CasbinService struct{}

func NewCasbinService() *CasbinService {
	return &CasbinService{}
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// getCasbinModel 获取Casbin模型（支持角色继承）
func getCasbinModel() (model.Model, error) {
	text := `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _
	g2 = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && keyMatch2(r.obj, p.obj) && (r.act == p.act || p.act == '*')
	`

	// 如果启用角色继承，使用更复杂的匹配器
	if global.TD27_CONFIG.Casbin.EnableRoleHierarchy {
		text = `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _
		g2 = _, _

		[policy_effect]
			e = some(where (p.eft == allow))

		[matchers]
			m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && (r.act == p.act || p.act == '*')
		`
	}

	return model.NewModelFromString(text)
}

func (cs *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		// 使用基于统一权限表的适配器
		adapter := casbinpkg.NewPermissionAdapter(global.TD27_DB)

		m, err := getCasbinModel()
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}

		var initErr error
		syncedCachedEnforcer, initErr = casbin.NewSyncedCachedEnforcer(m, adapter)
		if initErr != nil {
			zap.L().Error("Casbin enforcer初始化失败!", zap.Error(initErr))
			return
		}

		// 配置缓存TTL
		cacheTTL := global.TD27_CONFIG.Casbin.CacheTTL
		if cacheTTL <= 0 {
			cacheTTL = 3600 // 默认1小时
		}
		syncedCachedEnforcer.SetExpireTime(time.Duration(cacheTTL) * time.Second)

		// 启用自动加载策略
		if global.TD27_CONFIG.Casbin.AutoLoadInterval > 0 {
			syncedCachedEnforcer.StartAutoLoadPolicy(time.Duration(global.TD27_CONFIG.Casbin.AutoLoadInterval) * time.Second)
		}

		if err = syncedCachedEnforcer.LoadPolicy(); err != nil {
			zap.L().Error("Casbin策略加载失败!", zap.Error(err))
			return
		}

		global.TD27_LOG.Info("Casbin enforcer初始化完成，策略已加载",
			zap.Bool("roleHierarchy", global.TD27_CONFIG.Casbin.EnableRoleHierarchy),
			zap.Int("cacheTTL", cacheTTL))
	})
	return syncedCachedEnforcer
}

// Enforce 执行权限检查（支持多角色）
func (cs *CasbinService) Enforce(roleIDs []uint, path string, method string) (bool, error) {
	e := cs.Casbin()

	// 尝试所有角色，只要有一个通过就允许
	for _, roleID := range roleIDs {
		sub := strconv.Itoa(int(roleID))

		zap.L().Info("Enforce debug",
			zap.String("sub", sub),
			zap.String("obj", path),
			zap.String("act", string(modelSysManagement.HTTPMethodToAction(method))),
		)

		success, err := e.Enforce(sub, path, modelSysManagement.HTTPMethodToAction(method).String())
		if err != nil {
			return false, err
		}
		if success {
			return true, nil
		}
	}

	return false, nil
}

// ReloadPolicy 重新加载策略
func (cs *CasbinService) ReloadPolicy() error {
	return cs.Casbin().LoadPolicy()
}

// RebuildRolePolicies 更新角色的API权限策略（先删除旧策略，再添加新策略）
func (cs *CasbinService) RebuildRolePolicies(roleID uint, permissions []modelSysManagement.PermissionModel) error {
	e := cs.Casbin()
	if e == nil {
		return errors.New("casbin enforcer not initialized")
	}
	roleIDStr := strconv.Itoa(int(roleID))

	// 删除该角色的所有策略（fieldIndex 0 = sub/subject/roleID）
	_, err := e.RemoveFilteredPolicy(0, roleIDStr)
	if err != nil {
		return fmt.Errorf("remove old policies failed: %w", err)
	}

	// 如果没有新权限，直接返回
	if len(permissions) == 0 {
		return nil
	}

	// 构建新策略列表
	newPolicies := make([][]string, 0, len(permissions))
	for _, perm := range permissions {
		newPolicies = append(newPolicies, []string{
			roleIDStr,
			perm.Resource,
			string(perm.Action),
		})
	}

	// 批量添加新策略
	_, err = e.AddPolicies(newPolicies)
	if err != nil {
		return fmt.Errorf("add new policies failed: %w", err)
	}

	return nil
}

// AddRoleInheritance 添加角色继承关系
func (cs *CasbinService) AddRoleInheritance(childRoleID, parentRoleID uint) error {
	if !global.TD27_CONFIG.Casbin.EnableRoleHierarchy {
		return errors.New("role hierarchy is disabled")
	}

	child := strconv.Itoa(int(childRoleID))
	parent := strconv.Itoa(int(parentRoleID))

	e := cs.Casbin()
	_, err := e.AddGroupingPolicy(child, parent)
	return err
}

// RemoveRoleInheritance 移除角色继承关系
func (cs *CasbinService) RemoveRoleInheritance(childRoleID, parentRoleID uint) error {
	child := strconv.Itoa(int(childRoleID))
	parent := strconv.Itoa(int(parentRoleID))

	e := cs.Casbin()
	_, err := e.RemoveGroupingPolicy(child, parent)
	return err
}

// GetInheritedRoles 获取角色的所有继承角色
func (cs *CasbinService) GetInheritedRoles(roleID uint) ([]string, error) {
	sub := strconv.Itoa(int(roleID))
	e := cs.Casbin()
	return e.GetRolesForUser(sub)
}

// EnforceSubject 检查单个subject是否有权限（用于服务令牌）
func (cs *CasbinService) EnforceSubject(subject, path, method string) (bool, error) {
	e := cs.Casbin()
	return e.Enforce(subject, path, modelSysManagement.HTTPMethodToAction(method).String())
}

// RebuildSubjectPolicies 重建subject的所有策略（用于服务令牌）
func (cs *CasbinService) RebuildSubjectPolicies(subject string, policies [][]string) error {
	e := cs.Casbin()
	if e == nil {
		return errors.New("casbin enforcer not initialized")
	}

	global.TD27_LOG.Info("RebuildSubjectPolicies",
		zap.String("subject", subject),
		zap.Int("policyCount", len(policies)),
		zap.Any("policies", policies))

	// 删除该subject的所有现有策略
	_, err := e.RemoveFilteredPolicy(0, subject)
	if err != nil {
		return fmt.Errorf("remove subject policies failed: %w", err)
	}

	// 添加新策略
	if len(policies) > 0 {
		added, err := e.AddPolicies(policies)
		if err != nil {
			return fmt.Errorf("add subject policies failed: %w", err)
		}
		global.TD27_LOG.Info("AddPolicies result", zap.Bool("added", added))

		// 验证
		existing, _ := e.GetFilteredPolicy(0, subject)
		global.TD27_LOG.Info("After add", zap.Int("existingCount", len(existing)))
	}

	return nil
}

// RemoveSubjectPolicies 移除subject的所有策略
func (cs *CasbinService) RemoveSubjectPolicies(subject string) error {
	e := cs.Casbin()
	if e == nil {
		return errors.New("casbin enforcer not initialized")
	}

	_, err := e.RemoveFilteredPolicy(0, subject)
	if err != nil {
		return fmt.Errorf("remove subject policies failed: %w", err)
	}
	return nil
}

// RemoveResourcePolicy 移除指定资源的所有策略
func (cs *CasbinService) RemoveResourcePolicy(resource, action string) error {
	e := cs.Casbin()
	if e == nil {
		return errors.New("casbin enforcer not initialized")
	}

	_, err := e.RemoveFilteredPolicy(1, resource, action)
	if err != nil {
		return fmt.Errorf("remove resource policy failed: %w", err)
	}
	return nil
}

// UpdateResourcePolicies 批量更新指定资源的策略（用于API路径/方法变更时同步角色和服务令牌策略）
func (cs *CasbinService) UpdateResourcePolicies(oldResource, oldAction, newResource, newAction string) error {
	e := cs.Casbin()
	if e == nil {
		return errors.New("casbin enforcer not initialized")
	}

	// 获取所有匹配旧资源+动作的策略
	oldPolicies, err := e.GetFilteredPolicy(1, oldResource, oldAction)
	if err != nil {
		return fmt.Errorf("get old policies failed: %w", err)
	}

	global.TD27_LOG.Info("UpdateResourcePolicies",
		zap.String("oldResource", oldResource),
		zap.String("oldAction", oldAction),
		zap.String("newResource", newResource),
		zap.String("newAction", newAction),
		zap.Int("matchedPolicies", len(oldPolicies)))

	if len(oldPolicies) > 0 {
		// 删除旧策略
		removeBool, err := e.RemoveFilteredPolicy(1, oldResource, oldAction)
		if err != nil {
			return fmt.Errorf("remove old policies failed: %w", err)
		}
		global.TD27_LOG.Debug("RemovePolicy result", zap.Bool("removed", removeBool))
	}

	// 构建新策略列表（只更新obj和act，保留sub）
	newPolicies := make([][]string, 0, len(oldPolicies))
	for _, p := range oldPolicies {
		newPolicies = append(newPolicies, []string{p[0], newResource, newAction})
	}

	// 并添加新策略
	addBool, err := e.AddPolicies(newPolicies)
	if err != nil {
		return fmt.Errorf("add new policies failed: %w", err)
	}
	global.TD27_LOG.Debug("AddPolicies result", zap.Bool("added", addBool))

	return nil
}
