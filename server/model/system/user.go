package system

import (
	uuid "github.com/satori/go.uuid"
	"server/global"
)

type UserModel struct {
	global.TD27_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`           // 用户UUID
	Username    string    `json:"username" gorm:"index;unique;comment:用户登录名"` // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"` // 用户手机号
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`  // 用户邮箱
	Active      bool      `json:"active"`                      // 是否活跃
	RoleModelID uint      `json:"roleID"`                      // 用户角色
}

func (UserModel) TableName() string {
	return "sys_user"
}

type UserResult struct {
	ID          uint
	UUID        uuid.UUID `json:"uuid"`     // 用户UUID
	Username    string    `json:"username"` // 用户登录名
	Phone       string    `json:"phone"`    // 用户手机号
	Email       string    `json:"email"`    // 用户邮箱
	Active      bool      `json:"active"`   // 是否活跃
	RoleModelID uint      `json:"roleID"`   // 用户角色
	RoleName    string    `json:"role"`
}
