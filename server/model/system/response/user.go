package response

import "server/model/system"

type LoginResponse struct {
	User      system.UserModel `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
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
