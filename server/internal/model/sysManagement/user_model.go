package sysManagement

import (
	"server/internal/model/common"
)

// UserModel 用户模型（支持多角色）
type UserModel struct {
	common.Td27Model
	Username string `json:"username" gorm:"unique;size:191;comment:用户名" binding:"required"` // 用户名
	Password string `json:"-"  gorm:"not null;comment:密码"`
	Phone    string `json:"phone"  gorm:"comment:手机号"`                          // 手机号
	Email    string `json:"email"  gorm:"comment:邮箱" binding:"omitempty,email"` // 邮箱
	Active   bool   `json:"active"`                                             // 是否活跃
	DeptID   uint   `json:"deptId" gorm:"column:dept_id;comment:部门ID"`          // 部门ID（用于数据权限）
	// 多角色支持 - 替代原来的 RoleID uint
	Roles []*RoleModel `json:"roles" gorm:"many2many:sys_management_user_roles;joinForeignKey:user_id;joinReferences:role_id"`
	//RoleID uint `json:"roleId" gorm:"-"` // 虚拟字段，用于兼容旧API
}

func (UserModel) TableName() string {
	return "sys_management_user"
}

// GetPrimaryRoleID 获取主角色ID（兼容旧代码）
func (u *UserModel) GetPrimaryRoleID() uint {
	if len(u.Roles) > 0 {
		return u.Roles[0].ID
	}
	return 0
}

// GetAllRoleIDs 获取所有角色ID
func (u *UserModel) GetAllRoleIDs() []uint {
	ids := make([]uint, 0, len(u.Roles))
	for _, role := range u.Roles {
		ids = append(ids, role.ID)
	}
	return ids
}

// HasRole 检查用户是否有指定角色
func (u *UserModel) HasRole(roleID uint) bool {
	for _, role := range u.Roles {
		if role.ID == roleID {
			return true
		}
	}
	return false
}

// DeptModel 部门模型（用于数据权限）
type DeptModel struct {
	common.Td27Model
	DeptName string `json:"deptName" gorm:"size:100;not null;comment:部门名称"`
	ParentID uint   `json:"parentId" gorm:"index;comment:父部门ID"`
	Path     string `json:"path" gorm:"size:500;index;comment:部门路径(物化路径),如:/1/2/3/"`
	Sort     uint   `json:"sort" gorm:"default:0"`
	Status   bool   `json:"status" gorm:"default:true"`
}

// GetFullPath 获取完整路径（包含自己）
func (d *DeptModel) GetFullPath() string {
	if d.Path == "" {
		return "/" + string(rune(d.ID))
	}
	return d.Path + string(rune(d.ID)) + "/"
}

// IsAncestorOf 检查当前部门是否是目标部门的祖先
func (d *DeptModel) IsAncestorOf(targetPath string) bool {
	fullPath := d.GetFullPath()
	return len(targetPath) > len(fullPath) && targetPath[:len(fullPath)] == fullPath
}

// IsDescendantOf 检查当前部门是否是目标部门的后代
func (d *DeptModel) IsDescendantOf(ancestorPath string) bool {
	return len(d.Path) >= len(ancestorPath) && d.Path[:len(ancestorPath)] == ancestorPath
}

func (DeptModel) TableName() string {
	return "sys_management_dept"
}
