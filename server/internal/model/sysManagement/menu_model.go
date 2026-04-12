package sysManagement

import (
	"server/internal/model/common"
)

// MenuModel 菜单表（独立领域表）
type MenuModel struct {
	common.Td27Model
	MenuName   string `json:"menu_name" gorm:"unique;size:100;not null;comment:菜单名称"`
	Icon       string `json:"icon" gorm:"size:100;comment:图标"`
	Path       string `json:"path" gorm:"unique;size:200;comment:路由路径"`
	Component  string `json:"component" gorm:"size:200;comment:前端组件"`
	Redirect   string `json:"redirect" gorm:"size:200;comment:重定向"`
	ParentID   uint   `json:"parentId" gorm:"index;comment:父菜单ID"`
	Sort       uint   `json:"sort" gorm:"default:0;comment:排序"`
	Hidden     bool   `json:"hidden" gorm:"default:false;comment:是否隐藏"`
	KeepAlive  bool   `json:"keepAlive" gorm:"default:false;comment:缓存"`
	Affix      bool   `json:"affix" gorm:"default:false;comment:是否固定"`
	AlwaysShow bool   `json:"alwaysShow" gorm:"default:false;comment:一直显示根路由"`
	Title      string `json:"title" gorm:"unique;comment:菜单名"`
}

func (MenuModel) TableName() string {
	return "sys_management_menu"
}
