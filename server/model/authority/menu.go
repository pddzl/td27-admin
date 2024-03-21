package authority

import (
	"database/sql/driver"
	"encoding/json"
	"server/global"
)

type MenuModel struct {
	global.TD27_MODEL
	Pid       uint         `json:"pid"`                       // 父菜单ID
	Name      string       `json:"name,omitempty"`            // 路由名称
	Path      string       `json:"path" gorm:"unique"`        // 路由路径
	Redirect  string       `json:"redirect,omitempty"`        // 重定向
	Component string       `json:"component" gorm:"not null"` // 前端组件
	Sort      uint         `json:"sort" gorm:"not null"`      // 排序
	Meta      Meta         `json:"meta" gorm:"type:json"`     // 元数据
	Children  []MenuModel  `json:"children,omitempty" gorm:"-"`
	Roles     []*RoleModel `json:"-" gorm:"many2many:role_menus;"`
}

type Meta struct {
	Hidden     bool   `json:"hidden,omitempty"`  // 菜单是否隐藏
	Title      string `json:"title,omitempty"`   // 菜单名
	SvgIcon    string `json:"svgIcon,omitempty"` // svg图标
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
	return "authority_menu"
}
