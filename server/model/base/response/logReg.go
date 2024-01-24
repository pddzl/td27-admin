package response

import "server/model/authority"

type LoginResponse struct {
	User      authority.UserModel `json:"user"` // 用户
	Token     string              `json:"token"`
	ExpiresAt int64               `json:"expiresAt"` // 过期时间
}
