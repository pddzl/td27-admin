package system

import (
	"go.uber.org/zap"
	"server/global"
	systemModel "server/model/system"
	systemReq "server/model/system/request"
	"server/utils"
)

type MenuService struct{}

func getTreeMap(menuListFormat []systemModel.MenuModel, menuList []systemModel.MenuModel) {
	for index, menuF := range menuListFormat {
		for _, menu := range menuList {
			if menuF.ID == menu.Pid {
				// menuF 只是个复制值
				//menuF.Children = append(menuF.Children, menu)
				menuListFormat[index].Children = append(menuListFormat[index].Children, menu)
			}
		}
		if len(menuF.Children) > 0 {
			getTreeMap(menuF.Children, menuList)
		}
	}
}

func (ms *MenuService) GetMenus(roles []string) ([]systemModel.MenuModel, error) {
	var menuModels []systemModel.MenuModel
	err := global.TD27_DB.Preload("Roles").Find(&menuModels).Error

	menuList := make([]systemModel.MenuModel, 0)
	for _, menu := range menuModels {
		for _, menuRole := range menu.Roles {
			if utils.IsContain(roles, menuRole.RoleName) {
				menuList = append(menuList, menu)
				continue
			}
		}
	}

	menuListFormat := make([]systemModel.MenuModel, 0)
	for _, menu := range menuList {
		if menu.Pid == 0 {
			menuListFormat = append(menuListFormat, menu)
		}
	}

	getTreeMap(menuListFormat, menuList)

	return menuListFormat, err
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
