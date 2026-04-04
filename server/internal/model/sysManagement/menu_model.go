package sysManagement

import (
	"database/sql/driver"
	"encoding/json"

	"server/internal/model/common"
)

type MenuModel struct {
	common.Td27Model
	Pid       uint         `json:"pid"`                       // 父菜单 ID
	Name      string       `json:"name,omitempty"`            // 路由名称
	Path      string       `json:"path" gorm:"unique;size:191"`        // 路由路径
	Redirect  string       `json:"redirect,omitempty"`        // 重定向
	Component string       `json:"component" gorm:"not null"` // 前端组件
	Sort      uint         `json:"sort" gorm:"not null"`      // 排序
	Meta      Meta         `json:"meta" gorm:"type:json"`     // 元数据
	Children  []*MenuModel `json:"children,omitempty" gorm:"-"`
	// 权限通过 sys_management_role_permissions 关联，不再使用 sys_management_role_menus
}

type Meta struct {
	Hidden     bool   `json:"hidden,omitempty"`  // 菜单是否隐藏
	Title      string `json:"title,omitempty"`   // 菜单名
	SvgIcon    string `json:"svgIcon,omitempty"` // svg 图标
	ElIcon     string `json:"elIcon,omitempty"`  // element-plus图标
	Affix      bool   `json:"affix,omitempty"`   // 是否固定
	KeepAlive  bool   `json:"keepAlive,omitempty"`
	AlwaysShow bool   `json:"alwaysShow,omitempty"` // 是否一直显示根路由
}

func (m Meta) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *Meta) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), m)
}

func (MenuModel) TableName() string {
	return "sys_management_menu"
}
