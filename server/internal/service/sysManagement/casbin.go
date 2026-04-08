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
	"server/internal/model/sysManagement"
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

		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, adapter)

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

		_ = syncedCachedEnforcer.LoadPolicy()

		global.TD27_LOG.Info("Casbin enforcer initialized with unified permission table",
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
		success, err := e.Enforce(sub, path, method)
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

// UpdateRoleAPIPermissions 更新角色的API权限（通过统一权限表）
func (cs *CasbinService) UpdateRoleAPIPermissions(roleID uint, apiPermissionIDs []uint) error {
	db := global.TD27_DB

	// 开始事务
	tx := db.Begin()

	// 1. 删除该角色现有的API权限
	if err := tx.Exec(`
		DELETE FROM sys_management_role_permissions 
		WHERE role_id = ? AND permission_id IN (
			SELECT id FROM sys_management_permission WHERE type = 'api'
		)
	`, roleID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("clear existing API permissions failed: %w", err)
	}

	// 2. 添加新的API权限关联
	for _, permID := range apiPermissionIDs {
		rp := sysManagement.RolePermissionModel{
			RoleID:       roleID,
			PermissionID: permID,
			DataScope:    "all",
		}
		if err := tx.Create(&rp).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("add API permission failed: %w", err)
		}
	}

	// 3. 提交事务
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit transaction failed: %w", err)
	}

	// 4. 重新加载Casbin策略
	go func() {
		if err := cs.ReloadPolicy(); err != nil {
			global.TD27_LOG.Error("重新加载Casbin策略失败", zap.Error(err))
		}
	}()

	return nil
}
