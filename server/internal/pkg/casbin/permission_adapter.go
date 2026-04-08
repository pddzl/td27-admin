package casbin

import (
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"gorm.io/gorm"
)

// PermissionAdapter 基于统一权限表的Casbin适配器
type PermissionAdapter struct {
	db *gorm.DB
}

// NewPermissionAdapter 创建权限适配器
func NewPermissionAdapter(db *gorm.DB) *PermissionAdapter {
	return &PermissionAdapter{db: db}
}

// LoadPolicy 从统一权限表加载策略
func (a *PermissionAdapter) LoadPolicy(mod model.Model) error {
	// 查询所有API权限和角色关联
	var results []struct {
		RoleID   uint   `gorm:"column:role_id"`
		Resource string `gorm:"column:resource"`
		Method   string `gorm:"column:method"`
	}

	err := a.db.Raw(`
		SELECT rp.role_id, p.resource, p.method
		FROM sys_management_role_permissions rp
		JOIN sys_management_permission p ON rp.permission_id = p.id
		WHERE p.domain = 'api'
	`).Scan(&results).Error

	if err != nil {
		return fmt.Errorf("load policy from permission table failed: %w", err)
	}

	// 加载到Casbin模型
	for _, r := range results {
		roleID := strconv.Itoa(int(r.RoleID))
		line := fmt.Sprintf("p, %s, %s, %s", roleID, r.Resource, r.Method)
		persist.LoadPolicyLine(line, mod)
	}

	return nil
}

// SavePolicy 保存策略到统一权限表（不实现，通过API管理）
func (a *PermissionAdapter) SavePolicy(mod model.Model) error {
	// 策略通过统一权限表管理，不需要保存回Casbin格式
	return nil
}

// AddPolicy 添加策略（通过统一权限表管理，不直接操作）
func (a *PermissionAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	// 策略通过统一权限表管理
	return nil
}

// RemovePolicy 移除策略（通过统一权限表管理，不直接操作）
func (a *PermissionAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	// 策略通过统一权限表管理
	return nil
}

// RemoveFilteredPolicy 移除过滤策略（通过统一权限表管理，不直接操作）
func (a *PermissionAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	// 策略通过统一权限表管理
	return nil
}

// LoadFilteredPolicy 加载过滤策略
func (a *PermissionAdapter) LoadFilteredPolicy(mod model.Model, filter interface{}) error {
	return a.LoadPolicy(mod)
}

// IsFiltered 是否过滤
func (a *PermissionAdapter) IsFiltered() bool {
	return false
}

// AddPolicies 批量添加策略
func (a *PermissionAdapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	return nil
}

// RemovePolicies 批量移除策略
func (a *PermissionAdapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return nil
}

// UpdatePolicy 更新策略
func (a *PermissionAdapter) UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error {
	return nil
}

// UpdatePolicies 批量更新策略
func (a *PermissionAdapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	return nil
}

// UpdateFilteredPolicies 更新过滤策略
func (a *PermissionAdapter) UpdateFilteredPolicies(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) error {
	return nil
}
