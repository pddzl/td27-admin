package user

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type AddUserReq struct {
	Username    string `json:"username" binding:"required"`     // 用户名
	Password    string `json:"password" binding:"required"`     // 密码
	Phone       string `json:"phone" binding:"omitempty,phone"` // 手机号
	Email       string `json:"email" binding:"omitempty,email"` // 邮箱
	Active      bool   `json:"active"`                          // 是否活跃
	RoleModelID uint   `json:"roleId" binding:"required"`       // 角色 ID
}

// PhoneValidation 自定义手机号码校验函数
func PhoneValidation(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	re := regexp.MustCompile("^1[0-9]{10}")
	return re.MatchString(phone)
}

type UpdateUserReq struct {
	ID          uint   `json:"id" binding:"required"`           // 用户 ID
	Username    string `json:"username" binding:"required"`     // 用户名
	Phone       string `json:"phone" binding:"omitempty,phone"` // 手机号
	Email       string `json:"email" binding:"omitempty,email"` // 邮箱
	Active      bool   `json:"active"`                          // 是否活跃
	RoleModelID uint   `json:"roleId" binding:"required"`       // 角色 ID
}

type ModifyPasswdReq struct {
	ID          uint   `json:"id" binding:"required"`          // 用户 ID
	OldPassword string `json:"oldPassword" binding:"required"` // 旧密码
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

type SwitchActiveReq struct {
	ID     uint `json:"id" binding:"required"` // 用户 ID
	Active bool `json:"active"`                // 是否启用
}

type UserResp struct {
	UserModel
	RoleName string `json:"roleName"` // 角色名
}
