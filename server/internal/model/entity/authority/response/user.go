package response

import (
	"server/internal/model/entity/authority"
)

type UserResult struct {
	authority.UserModel
	RoleName string `json:"roleName"` // 角色名
}
