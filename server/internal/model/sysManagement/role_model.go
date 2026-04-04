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

// RolePermission 角色权限关联表（统一权限模型）
type RolePermission struct {
	RoleID       uint   `gorm:"column:role_id;primaryKey"`
	PermissionID uint   `gorm:"column:permission_id;primaryKey"`
	DataScope    string `gorm:"column:data_scope;size:20;default:'all';comment:数据权限范围:all全部|dept部门|self本人|custom自定义"`
	CustomSQL    string `gorm:"column:custom_sql;size:500;comment:自定义数据权限SQL条件"`
}

func (RolePermission) TableName() string {
	return "sys_management_role_permissions"
}
