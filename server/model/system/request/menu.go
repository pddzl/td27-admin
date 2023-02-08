package request

type Menu struct {
	Pid       uint   `json:"pid" validate:"required"`
	Name      string `json:"name"`
	Path      string `json:"path" validate:"required"`
	Redirect  string `json:"redirect"`
	Component string `json:"component"`
	Hidden    bool   `json:"hidden"` // 菜单是否隐藏
	Title     string `json:"title"`  // 菜单名
	Icon      string `json:"icon"`   // element图标
	Affix     bool   `json:"affix"`  // 是否固定
}
