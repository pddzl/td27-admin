package request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type AddUser struct {
	Username    string `json:"username" binding:"required"`     // 用户名
	Password    string `json:"password" binding:"required"`     // 密码
	Phone       string `json:"phone" binding:"omitempty,phone"` // 手机号
	Email       string `json:"email" binding:"omitempty,email"` // 邮箱
	Active      bool   `json:"active"`                          // 是否活跃
	RoleModelID uint   `json:"roleId" binding:"required"`       // 角色ID
}

// PhoneValidation 自定义手机号码校验函数
func PhoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	re := regexp.MustCompile("^1[0-9]{10}")
	return re.MatchString(phone)
}

type EditUser struct {
	ID          uint   `json:"id" binding:"required"`           // 用户ID
	Username    string `json:"username" binding:"required"`     // 用户名
	Phone       string `json:"phone" binding:"omitempty,phone"` // 手机号
	Email       string `json:"email" binding:"omitempty,email"` // 邮箱
	Active      bool   `json:"active"`                          // 是否活跃
	RoleModelID uint   `json:"roleId" binding:"required"`       // 角色ID
}

type ModifyPass struct {
	ID          uint   `json:"id" binding:"required"`          // 用户ID
	OldPassword string `json:"oldPassword" binding:"required"` // 旧密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

type SwitchActive struct {
	ID     uint `json:"id" binding:"required"` // 用户ID
	Active bool `json:"active"`                // 是否启用
}
