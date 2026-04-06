package sysManagement

import (
	"server/internal/model/common"
)

// RoleModel 角色模型（支持层级继承）
type RoleModel struct {
	common.Td27Model
	RoleName string       `json:"roleName" gorm:"unique;size:191" binding:"required"`
	ParentID *uint        `json:"parentId" gorm:"index;comment:父角色ID"` // 父角色ID，支持层级
	Parent   *RoleModel   `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children []*RoleModel `json:"children,omitempty" gorm:"foreignKey:ParentID"`
	// 权限通过 sys_management_role_permissions 关联
	// 权限缓存，避免频繁查询
	PermissionHash string `json:"-" gorm:"comment:权限哈希，用于缓存失效判断"`
}

func (RoleModel) TableName() string {
	return "sys_management_role"
}

// UserRole 用户角色关联表（支持多角色）
type UserRole struct {
	UserID uint `gorm:"column:user_id;primaryKey"`
	RoleID uint `gorm:"column:role_id;primaryKey"`
}

func (UserRole) TableName() string {
	return "sys_management_user_roles"
}
