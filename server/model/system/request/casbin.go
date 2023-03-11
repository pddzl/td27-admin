package request

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type ReqCasbin struct {
	RoleId      uint         `json:"roleId" validate:"required"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/base/captcha", Method: "POST"},
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/getUserInfo", Method: "GET"},
		{Path: "/menu/getMenus", Method: "GET"},
		{Path: "/user/editUser", Method: "POST"},
		{Path: "/user/modifyPass", Method: "POST"},
		{Path: "/jwt/joinInBlacklist", Method: "POST"},
	}
}
