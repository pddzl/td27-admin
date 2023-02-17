package system

import (
	"errors"
	"fmt"
	"server/global"
	systemModel "server/model/system"
	"server/utils"
)

type RoleService struct{}

func IsRole(username string, role string) bool {
	var userModel systemModel.UserModel
	global.TD27_DB.Where("username = ?", username).Preload("Roles").First(&userModel)

	var roles []string
	for _, value := range userModel.Roles {
		roles = append(roles, value.RoleName)
	}

	return utils.IsContain(roles, role)
}

func (rs *RoleService) GetRoles(username string) ([]systemModel.RoleModel, error) {
	if IsRole(username, "root") {
		var roleList []systemModel.RoleModel
		err := global.TD27_DB.Preload("Menus").Find(&roleList).Error

		return roleList, err
	}

	return nil, errors.New("没有权限")
}

func (rs *RoleService) AddRole(username string, roleName string) (*systemModel.RoleModel, error) {
	if IsRole(username, "root") {
		var roleModel systemModel.RoleModel
		roleModel.RoleName = roleName
		return &roleModel, global.TD27_DB.Create(&roleModel).Error
	}

	return &systemModel.RoleModel{}, errors.New("没有权限")
}

func (rs *RoleService) DeleteRole(id uint, username string) (err error) {
	if IsRole(username, "root") {
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

	return errors.New("没有权限")
}
