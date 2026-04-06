package sysManagement

// Menu 菜单请求结构（前端传入）
type Menu struct {
	ParentId   uint   `json:"parentId"`                     // 默认0 根目录
	MenuName   string `json:"menu_name"`                    // 名称（路由名）
	Path       string `json:"path" binding:"required"`      // 路径
	Redirect   string `json:"redirect"`                     // 重定向
	Component  string `json:"component" binding:"required"` // 前端组件
	Sort       uint   `json:"sort" binding:"required"`      // 排序
	Hidden     bool   `json:"hidden"`                       // 隐藏菜单
	Title      string `json:"title"`                        // 菜单名
	Icon       string `json:"icon"`                         // 图标
	Affix      bool   `json:"affix"`                        // 固定
	KeepAlive  bool   `json:"keepAlive"`                    // 缓存
	AlwaysShow bool   `json:"alwaysShow"`                   // 一直显示根路由
}

// UpdateMenuReq 更新菜单请求
type UpdateMenuReq struct {
	ID uint `json:"id" binding:"required"` // 菜单 ID
	Menu
}

// MenuResp 菜单响应（扁平化结构，直接返回给前端）
type MenuResp struct {
	ID         uint        `json:"id"`
	MenuName   string      `json:"menu_name"`
	Icon       string      `json:"icon"`
	Path       string      `json:"path"`
	Component  string      `json:"component"`
	Redirect   string      `json:"redirect"`
	ParentID   uint        `json:"parentId"`
	Sort       uint        `json:"sort"`
	Hidden     bool        `json:"hidden"`
	KeepAlive  bool        `json:"keepAlive"`
	Affix      bool        `json:"affix"`
	AlwaysShow bool        `json:"alwaysShow"`
	Title      string      `json:"title"`
	Children   []*MenuResp `json:"children,omitempty"`
}

// MenuElTreeResp el-tree菜单响应
type MenuElTreeResp struct {
	List    []MenuResp `json:"list"`
	MenuIds []uint     `json:"menuIds"`
}
