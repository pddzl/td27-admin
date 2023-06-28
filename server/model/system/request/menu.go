package request

type Menu struct {
	Pid       uint   `json:"pid"` // 默认0 根目录
	Name      string `json:"name"`
	Path      string `json:"path" validate:"required"`
	Redirect  string `json:"redirect"`
	Component string `json:"component" validate:"required"`
	Sort      uint   `json:"sort" validate:"required"`
	Meta      meta   `json:"meta"`
}

type meta struct {
	Hidden    bool   `json:"hidden"` // 菜单是否隐藏
	Title     string `json:"title"`  // 菜单名
	Icon      string `json:"icon"`   // element图标
	Affix     bool   `json:"affix"`  // 是否固定
	KeepAlive bool   `json:"keepAlive"`
}

type EditMenuReq struct {
	Id uint `json:"id" validate:"required"`
	Menu
}

type EditRoleMenu struct {
	RoleId uint   `json:"roleId"`
	Ids    []uint `json:"ids"`
}
