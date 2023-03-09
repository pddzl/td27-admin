package request

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type ReqCasbin struct {
	RoleId      uint         `json:"roleId" validate:"required"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
