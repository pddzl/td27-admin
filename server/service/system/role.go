package system

import (
	"errors"
	"server/global"
	systemModel "server/model/system"
	"server/utils"
)

type RoleService struct{}

func (rs *RoleService) GetRoles(username string) ([]systemModel.RoleModel, error) {
	var userModel systemModel.UserModel
	global.TD27_DB.Where("username = ?", username).Preload("Roles").First(&userModel)

	var roles []string
	for _, role := range userModel.Roles {
		roles = append(roles, role.RoleName)
	}

	if utils.IsContain(roles, "root") {
		var roleList []systemModel.RoleModel
		err := global.TD27_DB.Preload("Menus").Find(&roleList).Error

		return roleList, err
	}

	return nil, errors.New("没有权限")
}
