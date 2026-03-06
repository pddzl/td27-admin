package sysManagement

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type ReqCasbin struct {
	RoleId      uint         `json:"roleId" binding:"required"` // 角色 ID
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/captcha", Method: "POST"},
		{Path: "/login", Method: "POST"},
		{Path: "/logout", Method: "POST"},
		{Path: "/user/getUserInfo", Method: "GET"},
		{Path: "/user/update", Method: "POST"},
		{Path: "/user/modifyPasswd", Method: "POST"},
		{Path: "/menu/list", Method: "GET"},
	}
}
