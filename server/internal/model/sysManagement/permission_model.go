package sysManagement

import (
	"server/internal/model/common"
)

// PermissionType 权限类型
type PermissionType string

const (
	PermissionTypeMenu   PermissionType = "menu"
	PermissionTypeAPI    PermissionType = "api"
	PermissionTypeButton PermissionType = "button"
	PermissionTypeData   PermissionType = "data"
)

// PermissionModel 统一权限身份表（仅用于RBAC授权）
type PermissionModel struct {
	common.Td27Model
	Name     string         `json:"name" gorm:"size:100;not null;comment:权限名称"`
	Domain   PermissionType `gorm:"type:varchar(20);not null;check:domain IN ('menu','api','button','data')"`
	Resource string         `json:"resource" gorm:"size:200;not null;comment:资源标识，如:/api/user或menu:users"`
	Action   string         `json:"action" gorm:"size:20;default:'all';comment:操作:all|view|create|update|delete"`
	// 关联的领域表ID
	DomainID uint `json:"domainId" gorm:"index;comment:关联领域表ID(menu/api/button)"`
}

func (PermissionModel) TableName() string {
	return "sys_management_permission"
}

// PermissionIdentity 权限身份（用于Casbin）
type PermissionIdentity struct {
	Type     string `json:"type"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

// ToCasbinRule 转换为Casbin规则格式
func (p *PermissionModel) ToCasbinRule() (obj string, act string) {
	return p.Resource, p.Action
}
