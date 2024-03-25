package response

import "server/model/authority"

type UserResult struct {
	authority.UserModel
	RoleName string `json:"roleName"` // 角色名
}
