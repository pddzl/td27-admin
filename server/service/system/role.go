package system

import (
	"errors"
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

func (rs *RoleService) AddRole(username string, roleName string) (err error) {
	if IsRole(username, "root") {
		var roleModel systemModel.RoleModel
		roleModel.RoleName = roleName
		return global.TD27_DB.Create(&roleModel).Error
	}

	return errors.New("没有权限")
}
