package sysManagement

// ApiTreeResp API树响应
type ApiTreeResp struct {
	List       []*ApiTreeNode `json:"list"`
	CheckedKey []string       `json:"checkedKey"`
	CheckedIds []uint         `json:"checkedIds"`
}

// UpdateApiReq 更新API请求
type UpdateApiReq struct {
	ID       uint   `json:"id" binding:"required"`
	ApiName  string `json:"apiName" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	ApiGroup string `json:"apiGroup"`
}

// ListApiReq API列表请求
type ListApiReq struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Path     string `json:"path" form:"path"`
	Method   string `json:"method" form:"method"`
	ApiGroup string `json:"apiGroup" form:"apiGroup"`
}

// CreateApiReq 创建API请求
type CreateApiReq struct {
	ApiName  string `json:"apiName" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Method   string `json:"method" binding:"required"`
	ApiGroup string `json:"apiGroup"`
}

// ApiTreeNode API树节点
type ApiTreeNode struct {
	ID       uint           `json:"id"`
	ApiName  string         `json:"apiName"`
	Path     string         `json:"path"`
	Method   string         `json:"method"`
	ApiGroup string         `json:"apiGroup"`
	Children []*ApiTreeNode `json:"children,omitempty"`
}
