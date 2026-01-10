package response

import (
	"server/internal/model/authority/user"
)

type LoginResponse struct {
	User      user.UserModel `json:"user"` // 用户
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"` // 过期时间
}
