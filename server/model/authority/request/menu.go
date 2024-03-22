package request

type Menu struct {
	Pid       uint   `json:"pid"`                          // 默认0 根目录
	Name      string `json:"name"`                         // 名称
	Path      string `json:"path" binding:"required"`      // 路径
	Redirect  string `json:"redirect"`                     // 重定向
	Component string `json:"component" binding:"required"` // 前端组件
	Sort      uint   `json:"sort" binding:"required"`      // 排序
	Meta      meta   `json:"meta"`
}

type meta struct {
	Hidden     bool   `json:"hidden"`     // 隐藏菜单
	Title      string `json:"title"`      // 菜单名
	Icon       string `json:"icon"`       // element图标
	Affix      bool   `json:"affix"`      // 组件固定
	KeepAlive  bool   `json:"keepAlive"`  // 组件缓存
	AlwaysShow bool   `json:"alwaysShow"` // 是否一直显示根路由
}

type EditMenuReq struct {
	ID uint `json:"id" binding:"required"` // 菜单ID
	Menu
}

type EditRoleMenu struct {
	RoleId uint   `json:"roleId"` // 角色ID
	Ids    []uint `json:"ids"`    // 菜单ID
}
