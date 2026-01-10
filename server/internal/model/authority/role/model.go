package role

import (
	"server/internal/model/authority/menu"
	"server/internal/model/common"
)

type RoleModel struct {
	common.Td27Model
	RoleName string `json:"roleName" gorm:"unique" binding:"required"`
	//Users    []*UserModel `json:"users"`
	Menus []*menu.MenuModel `json:"menus" gorm:"many2many:role_menus;"`
}

func (RoleModel) TableName() string {
	return "authority_role"
}
