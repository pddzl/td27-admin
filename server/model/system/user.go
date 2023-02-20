package system

import (
	"server/global"
)

type UserModel struct {
	global.TD27_MODEL
	Username    string `json:"username" gorm:"index;unique;comment:用户登录名" validate:"required"` // 用户登录名
	Password    string `json:"-"  gorm:"comment:用户登录密码" validate:"required"`
	Phone       string `json:"phone"  gorm:"comment:用户手机号" validate:"required"` // 用户手机号
	Email       string `json:"email"  gorm:"comment:用户邮箱" validate:"required"`  // 用户邮箱
	Active      bool   `json:"active" validate:"required"`                      // 是否活跃
	RoleModelID uint   `json:"roleID" validate:"required"`                      // 用户角色
}

func (UserModel) TableName() string {
	return "sys_user"
}

type UserResult struct {
	ID          uint
	Username    string `json:"username"` // 用户登录名
	Phone       string `json:"phone"`    // 用户手机号
	Email       string `json:"email"`    // 用户邮箱
	Active      bool   `json:"active"`   // 是否活跃
	RoleModelID uint   `json:"roleID"`   // 用户角色
	RoleName    string `json:"role"`
}
