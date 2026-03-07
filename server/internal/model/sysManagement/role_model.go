package sysManagement

import (
	"server/internal/model/common"
)

type RoleModel struct {
	common.Td27Model
	RoleName string `json:"roleName" gorm:"unique" binding:"required"`
	//Users    []*UserModel `json:"users"`
	Menus []*MenuModel `json:"menus" gorm:"many2many:sys_management_role_menus;joinForeignKey:RoleID;joinReferences:MenuID"`
}

func (RoleModel) TableName() string {
	return "sys_management_role"
}
