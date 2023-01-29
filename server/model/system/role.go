package system

import "server/global"

type RoleModel struct {
	global.TD27_MODEL
	RoleName string `json:"roleName"`
}

func (RoleModel) TableName() string {
	return "sys_role"
}
