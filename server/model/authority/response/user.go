package response

import (
	"server/model/base"
	"time"
)

type LoginResponse struct {
	User      base.UserModel `json:"user"` // 用户
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"` // 过期时间
}

type UserResult struct {
	CreatedAt   time.Time `json:"createdAt"` // 创建时间
	ID          uint      // 用户ID
	Username    string    `json:"username"` // 用户名
	Phone       string    `json:"phone"`    // 手机号
	Email       string    `json:"email"`    // 邮箱
	Active      bool      `json:"active"`   // 是否活跃
	RoleModelID uint      `json:"roleId"`   // 角色ID
	RoleName    string    `json:"role"`     // 角色名
}
