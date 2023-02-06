package system

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type MenuModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`  // 主键ID
	CreatedAt time.Time      `json:"-"`                     // 创建时间
	UpdatedAt time.Time      `json:"-"`                     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`        // 删除时间
	Pid       uint           `json:"pid"`                   // 父菜单ID
	Name      string         `json:"name"`                  // 路由名称
	Path      string         `json:"path"`                  // 路由路径
	Redirect  string         `json:"redirect,omitempty"`    // 重定向
	Component string         `json:"component"`             // 对端组件
	Meta      Meta           `json:"meta" gorm:"type:json"` // 元数据
	Roles     []*RoleModel   `json:"-" gorm:"many2many:role_menus;"`
}

type Meta struct {
	Hidden  bool   `json:"hidden"`  // 菜单是否隐藏
	Title   string `json:"title"`   // 菜单名
	SvgIcon string `json:"svgIcon"` // 图标
	Affix   bool   `json:"affix"`   // 是否固定
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
