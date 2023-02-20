package system

import (
	"server/global"
)

type UserModel struct {
	global.TD27_MODEL
	Username    string `json:"username" gorm:"index;unique;comment:用户名"` // 用户名
	Password    string `json:"-"  gorm:"comment:密码"`
	Phone       string `json:"phone"  gorm:"comment:手机号"` // 手机号
	Email       string `json:"email"  gorm:"comment:邮箱"`  // 邮箱
	Active      bool   `json:"active"`                    // 是否活跃
	RoleModelID uint   `json:"roleID"`                    // 角色ID
}

func (UserModel) TableName() string {
	return "sys_user"
}

type UserResult struct {
	ID          uint
	Username    string `json:"username"` // 用户名
	Phone       string `json:"phone"`    // 手机号
	Email       string `json:"email"`    // 邮箱
	Active      bool   `json:"active"`   // 是否活跃
	RoleModelID uint   `json:"roleID"`   // 角色ID
	RoleName    string `json:"role"`     // 角色名
}

type AddUser struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
	//Phone       string `json:"phone" validate:"required,regexp=^[1][0-9]{10}$"` // 手机号
	Phone       string `json:"phone" validate:"required"`       // 手机号
	Email       string `json:"email" validate:"required,email"` // 邮箱
	Active      bool   `json:"active"`                          // 是否活跃
	RoleModelID uint   `json:"roleID" validate:"required"`      // 角色ID
}
