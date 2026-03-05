package sysManagement

import (
	"server/internal/model/common"
)

type UserModel struct {
	common.Td27Model
	Username string `json:"username" gorm:"unique;comment:用户名" binding:"required"` // 用户名
	Password string `json:"-"  gorm:"not null;comment:密码"`
	Phone    string `json:"phone"  gorm:"comment:手机号"`                                // 手机号
	Email    string `json:"email"  gorm:"comment:邮箱" binding:"omitempty,email"`       // 邮箱
	Active   bool   `json:"active"`                                                   // 是否活跃
	RoleID   uint   `json:"roleId" gorm:"column:role_id;not null" binding:"required"` // 角色 ID
}

func (UserModel) TableName() string {
	return "authority_user"
}
