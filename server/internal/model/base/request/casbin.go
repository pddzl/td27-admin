package request

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type ReqCasbin struct {
	RoleId      uint         `json:"roleId" binding:"required"` // 角色ID
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/logReg/captcha", Method: "POST"},
		{Path: "/logReg/login", Method: "POST"},
		{Path: "/logReg/logout", Method: "POST"},
		{Path: "/user/getUserInfo", Method: "GET"},
		{Path: "/user/editUser", Method: "POST"},
		{Path: "/user/modifyPass", Method: "POST"},
		{Path: "/menu/getMenus", Method: "GET"},
	}
}
