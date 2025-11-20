package response

import (
	"server/internal/model/entity/authority"
)

type LoginResponse struct {
	User      authority.UserModel `json:"user"` // 用户
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"` // 过期时间
}
