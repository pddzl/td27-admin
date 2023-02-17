package system

import "server/global"

type RoleModel struct {
	global.TD27_MODEL
	RoleName string       `json:"roleName" gorm:"unique"`
	Menus    []*MenuModel `json:"menus" gorm:"many2many:role_menus;"`
}

func (RoleModel) TableName() string {
	return "sys_role"
}
