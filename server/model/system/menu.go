package system

import (
	"database/sql/driver"
	"encoding/json"
	"server/global"
)

type MenuModel struct {
	global.TD27_MODEL
	Pid       uint         `json:"pid"`                   // 父菜单ID
	Name      string       `json:"name,omitempty"`        // 路由名称
	Path      string       `json:"path"`                  // 路由路径
	Redirect  string       `json:"redirect,omitempty"`    // 重定向
	Component string       `json:"component"`             // 前端组件
	Meta      Meta         `json:"meta" gorm:"type:json"` // 元数据
	Children  []MenuModel  `json:"children,omitempty" gorm:"-"`
	Roles     []*RoleModel `json:"-" gorm:"many2many:role_menus;"`
}

type Meta struct {
	Hidden  bool   `json:"hidden,omitempty"`  // 菜单是否隐藏
	Title   string `json:"title,omitempty"`   // 菜单名
	Icon    string `json:"icon,omitempty"`    // 图标
	ElIcon  string `json:"elIcon,omitempty"`  // element图标
	SvgIcon string `json:"svgIcon,omitempty"` // svg图标
	Affix   bool   `json:"affix,omitempty"`   // 是否固定
}

func (m Meta) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	return string(b), err
}

func (m *Meta) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), m)
}

func (MenuModel) TableName() string {
	return "sys_menu"
}
