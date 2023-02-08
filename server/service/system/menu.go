package system

import (
	"go.uber.org/zap"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	"server/utils"
)

type MenuService struct{}

func (ms *MenuService) GetMenus(roles []string) (menuList []systemModel.MenuModel, err error) {
	var menuModels []systemModel.MenuModel
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

func (ms *MenuService) AddMenu(menuRaw systemReq.Menu) bool {
	var menuModel systemModel.MenuModel
	menuModel.Name = menuRaw.Name
	menuModel.Path = menuRaw.Path
	menuModel.Component = menuRaw.Component
	menuModel.Redirect = menuRaw.Redirect
	menuModel.Pid = menuRaw.Pid
	menuModel.Meta.Title = menuRaw.Title
	menuModel.Meta.Icon = menuRaw.Icon
	menuModel.Meta.Hidden = menuRaw.Hidden
	menuModel.Meta.Affix = menuRaw.Affix

	if err := global.TD27_DB.Create(&menuModel).Error; err != nil {
		global.TD27_LOG.Error("创建menu失败", zap.Error(err))
		return false
	}

	return true
}
