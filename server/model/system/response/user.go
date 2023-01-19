package response

import "server/model/system"

type LoginResponse struct {
	User      system.UserModel `json:"user"`
	Token     string           `json:"token"`
	ExpiresAt int64            `json:"expiresAt"`
}
