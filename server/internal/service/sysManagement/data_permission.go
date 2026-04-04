package sysManagement

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"

	"server/internal/global"
	"server/internal/model/sysManagement"
)



// DataPermissionService 数据权限服务
type DataPermissionService struct {
	db    *gorm.DB
	cache *sync.Map // 简单的内存缓存
}

// NewDataPermissionService 创建数据权限服务
func NewDataPermissionService() *DataPermissionService {
	return &DataPermissionService{
		db:    global.TD27_DB,
		cache: &sync.Map{},
	}
}

// GetUserDataPermission 获取用户的数据权限配置
// 如果用户有多个角色，取最宽松的权限（all > dept > self）
func (s *DataPermissionService) GetUserDataPermission(ctx context.Context, userID uint, resource string) (*sysManagement.DataPermission, error) {
	// 缓存key
	cacheKey := fmt.Sprintf("data_perm:%d:%s", userID, resource)

	// 尝试从缓存获取
	if cached, ok := s.cache.Load(cacheKey); ok {
		if perm, ok := cached.(*sysManagement.DataPermission); ok {
			return perm, nil
		}
	}

	// 查询用户的所有角色及其数据权限
	var results []struct {
		RoleID    uint   `gorm:"column:role_id"`
		DataScope string `gorm:"column:data_scope"`
		CustomSQL string `gorm:"column:custom_sql"`
	}

	err := s.db.WithContext(ctx).Raw(`
		SELECT rp.role_id, rp.data_scope, rp.custom_sql
		FROM sys_management_role_permissions rp
		JOIN sys_management_user_roles ur ON rp.role_id = ur.role_id
		JOIN sys_management_permission p ON rp.permission_id = p.id
		WHERE ur.user_id = ? AND p.resource = ? AND p.type = 'data'
	`, userID, resource).Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("query data permission failed: %w", err)
	}

	// 如果没有找到数据权限配置，默认只能看自己的数据
	if len(results) == 0 {
		perm := &sysManagement.DataPermission{
			Scope:  "self",
			UserID: userID,
		}
		s.cache.Store(cacheKey, perm)
		return perm, nil
	}

	// 获取用户部门ID
	var user sysManagement.UserModel
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("get user dept failed: %w", err)
	}

	// 选择最宽松的权限
	perm := s.selectMostPermissive(results, userID, user.DeptID)

	// 缓存结果（5分钟）
	s.cache.Store(cacheKey, perm)
	go func() {
		time.Sleep(5 * time.Minute)
		s.cache.Delete(cacheKey)
	}()

	return perm, nil
}

// selectMostPermissive 选择最宽松的数据权限
func (s *DataPermissionService) selectMostPermissive(results []struct {
	RoleID    uint   `gorm:"column:role_id"`
	DataScope string `gorm:"column:data_scope"`
	CustomSQL string `gorm:"column:custom_sql"`
}, userID, deptID uint) *sysManagement.DataPermission {
	// 权限优先级：all > dept > self > custom
	scopePriority := map[string]int{
		"all":    4,
		"dept":   3,
		"self":   2,
		"custom": 1,
	}

	var bestResult *struct {
		RoleID    uint   `gorm:"column:role_id"`
		DataScope string `gorm:"column:data_scope"`
		CustomSQL string `gorm:"column:custom_sql"`
	}
	bestPriority := 0

	for i := range results {
		priority := scopePriority[results[i].DataScope]
		if priority > bestPriority {
			bestPriority = priority
			bestResult = &results[i]
		}
	}

	if bestResult == nil {
		return &sysManagement.DataPermission{
			Scope:  "self",
			UserID: userID,
		}
	}

	return &sysManagement.DataPermission{
		Scope:     sysManagement.DataScope(bestResult.DataScope),
		DeptID:    deptID,
		UserID:    userID,
		CustomSQL: bestResult.CustomSQL,
	}
}

// ApplyDataScope 应用数据权限到查询
func (s *DataPermissionService) ApplyDataScope(db *gorm.DB, perm *sysManagement.DataPermission, tableAlias string) *gorm.DB {
	if tableAlias != "" {
		tableAlias = tableAlias + "."
	}

	switch perm.Scope {
	case sysManagement.DataScopeAll:
		// 全部数据，不添加过滤
		return db
	case sysManagement.DataScopeDept:
		// 本部门数据
		return db.Where(tableAlias+"dept_id = ?", perm.DeptID)
	case sysManagement.DataScopeSelf:
		// 仅本人数据
		return db.Where(tableAlias+"id = ?", perm.UserID)
	case sysManagement.DataScopeCustom:
		// 自定义SQL
		if perm.CustomSQL != "" {
			return db.Where(perm.CustomSQL)
		}
		return db
	default:
		// 默认只能看自己的
		return db.Where(tableAlias+"id = ?", perm.UserID)
	}
}

// ClearCache 清除权限缓存
func (s *DataPermissionService) ClearCache(userID uint) {
	// 清除该用户的所有缓存
	prefix := fmt.Sprintf("data_perm:%d:", userID)
	s.cache.Range(func(key, value interface{}) bool {
		if k, ok := key.(string); ok {
			if len(k) >= len(prefix) && k[:len(prefix)] == prefix {
				s.cache.Delete(key)
			}
		}
		return true
	})
}
