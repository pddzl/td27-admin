package system

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
		if len(menuListFormat[index].Children) > 0 {
			getTreeMap(menuListFormat[index].Children, menuList)
		}
	}
}

func (ms *MenuService) GetMenus(roles []string) ([]systemModel.MenuModel, error) {
	var menuModels []systemModel.MenuModel
	err := global.TD27_DB.Preload("Roles").Find(&menuModels).Error

	// 过滤角色拥有的路由
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

func (ms *MenuService) UpdateMenu(menuRaw systemReq.EditMenuReq) (err error) {
	var menuModel systemModel.MenuModel
	var metaData systemModel.Meta

	if errors.Is(global.TD27_DB.Where("id = ?", menuRaw.Id).First(&menuModel).Error, gorm.ErrRecordNotFound) {
		return errors.New("菜单不存在")
	}

	metaData.Icon = menuRaw.Icon
	metaData.Title = menuRaw.Title
	metaData.Hidden = menuRaw.Hidden
	metaData.Affix = menuRaw.Affix

	err = global.TD27_DB.Model(&menuModel).Updates(map[string]interface{}{"pid": menuRaw.Pid,
		"name":      menuRaw.Name,
		"path":      menuRaw.Path,
		"component": menuRaw.Component,
		"redirect":  menuRaw.Redirect,
		"meta":      metaData,
	}).Error

	return
}

func (ms *MenuService) DeleteMenu(id uint) (err error) {
	return global.TD27_DB.Unscoped().Delete(&systemModel.MenuModel{}, id).Error
}
