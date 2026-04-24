package casbin

import (
	"fmt"
	"log/slog"

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
	type policyRow struct {
		Sub      string `gorm:"column:sub"`
		Resource string `gorm:"column:resource"`
		Action   string `gorm:"column:action"`
	}

	var rows []policyRow

	// Merge role + token into ONE query
	err := a.db.Raw(`
		SELECT 
			CAST(rp.role_id AS TEXT) AS sub,
			p.resource,
			p.action
		FROM sys_management_role_permissions rp
		JOIN sys_management_permission p 
			ON rp.permission_id = p.id
		WHERE p.domain = 'api'

		UNION ALL

		SELECT 
			CONCAT('token:', tp.token_id) AS sub,
			p.resource,
			p.action
		FROM sys_tool_token_permission tp
		JOIN sys_management_permission p 
			ON tp.permission_id = p.id
		WHERE p.domain = 'api'
	`).Scan(&rows).Error

	if err != nil {
		return fmt.Errorf("load policy failed: %w", err)
	}

	// Single loop
	for _, r := range rows {
		line := fmt.Sprintf("p, %s, %s, %s", r.Sub, r.Resource, r.Action)
		persist.LoadPolicyLine(line, mod)
	}

	// Accurate logging
	slog.Info("Casbin策略加载完成",
		"policyCount", len(rows),
	)

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
