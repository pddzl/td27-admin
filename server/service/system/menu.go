package system

import (
	"server/global"
	"server/model/system"
	"server/utils"
)

type MenuService struct{}

func (ms *MenuService) GetMenus(roles []string) (menuList []system.MenuModel, err error) {
	var menuModels []system.MenuModel
	err = global.TD27_DB.Preload("Roles").Find(&menuModels).Error

	for _, menu := range menuModels {
		for _, menuRole := range menu.Roles {
			if utils.IsContain(roles, menuRole.RoleName) {
				menuList = append(menuList, menu)
				continue
			}
		}
	}

	return menuList, err
}
