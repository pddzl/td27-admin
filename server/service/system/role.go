package system

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
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

// DeleteRole 删除角色
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

// EditRole 编辑用户
func (rs *RoleService) EditRole(eRole systemReq.EditRole) (err error) {
	var roleModel systemModel.RoleModel
	err = global.TD27_DB.Where("id = ?", eRole.ID).First(&roleModel).Error
	if err != nil {
		global.TD27_LOG.Error("查询角色", zap.Error(err))
	}

	return global.TD27_DB.Model(&roleModel).Update("role_name", eRole.RoleName).Error
}

func getTreeMap1(menuListFormat []systemModel.MenuModel, menuList []*systemModel.MenuModel) {
	for index, menuF := range menuListFormat {
		for _, menu := range menuList {
			if menuF.ID == menu.Pid {
				menuListFormat[index].Children = append(menuListFormat[index].Children, *menu)
			}
		}
		if len(menuListFormat[index].Children) > 0 {
			getTreeMap1(menuListFormat[index].Children, menuList)
		}
	}
}

// GetRoleMenus 查找角色的menus
func (rs *RoleService) GetRoleMenus(id uint) ([]systemModel.MenuModel, error) {
	var roleModel systemModel.RoleModel
	err := global.TD27_DB.Where("id = ?", id).Preload("Menus").First(&roleModel).Error
	if err != nil {
		global.TD27_LOG.Error("GetRoleMenus 查找角色", zap.Error(err))
		return nil, err
	}

	// 获取menus树形结构
	menuListFormat := make([]systemModel.MenuModel, 0)
	for _, menu := range roleModel.Menus {
		if menu.Pid == 0 {
			menuListFormat = append(menuListFormat, *menu)
		}
	}

	getTreeMap1(menuListFormat, roleModel.Menus)

	return menuListFormat, nil
}
