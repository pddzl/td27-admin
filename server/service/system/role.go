package system

import (
	"fmt"
	"server/global"
	systemModel "server/model/system"
)

type RoleService struct{}

func (rs *RoleService) GetRoles() ([]systemModel.RoleModel, error) {
	var roleList []systemModel.RoleModel
	err := global.TD27_DB.Preload("Menus").Find(&roleList).Error

	return roleList, err
}

func (rs *RoleService) AddRole(roleName string) (*systemModel.RoleModel, error) {
	var roleModel systemModel.RoleModel
	roleModel.RoleName = roleName
	return &roleModel, global.TD27_DB.Create(&roleModel).Error

}

func (rs *RoleService) DeleteRole(id uint) (err error) {
	var roleModel systemModel.RoleModel

	err = global.TD27_DB.Where("id = ?", id).First(&roleModel).Error
	if err != nil {
		return fmt.Errorf("查询role -> %v", err)
	}

	err = global.TD27_DB.Unscoped().Delete(&roleModel).Error
	if err != nil {
		return fmt.Errorf("删除role -> %v", err)
	}

	// 清空关联
	err = global.TD27_DB.Model(&roleModel).Association("Menus").Clear()
	if err != nil {
		return fmt.Errorf("删除role关联menus -> %v", err)
	}

	return
}
