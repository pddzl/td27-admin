package system

import (
	uuid "github.com/satori/go.uuid"
	"server/global"
)

type UserModel struct {
	global.TD27_MODEL
	UUID     uuid.UUID   `json:"uuid" gorm:"index;comment:用户UUID"`    // 用户UUID
	Username string      `json:"userName" gorm:"index;comment:用户登录名"` // 用户登录名
	Password string      `json:"-"  gorm:"comment:用户登录密码"`
	Phone    string      `json:"phone"  gorm:"comment:用户手机号"`      // 用户手机号
	Email    string      `json:"email"  gorm:"comment:用户邮箱"`       // 用户邮箱
	Roles    []RoleModel `json:"roles" gorm:"many2many:user_role"` // 角色
}

func (UserModel) TableName() string {
	return "sys_user"
}
