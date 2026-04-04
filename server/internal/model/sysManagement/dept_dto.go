package sysManagement

// Dept 部门信息
type Dept struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	DeptName string `json:"deptName" gorm:"comment:部门名称" binding:"required"`
	ParentID uint   `json:"parentId" gorm:"comment:父部门ID"`
	Path     string `json:"path" gorm:"comment:部门路径(物化路径)"`
	Sort     uint   `json:"sort" gorm:"comment:排序"`
	Status   bool   `json:"status" gorm:"comment:状态"`
}

// CreateDeptReq 创建部门请求
type CreateDeptReq struct {
	DeptName string `json:"deptName" binding:"required"`
	ParentID uint   `json:"parentId"`
	Path     string `json:"path"`
	Sort     uint   `json:"sort"`
	Status   bool   `json:"status"`
}

// UpdateDeptReq 更新部门请求
type UpdateDeptReq struct {
	ID       uint   `json:"id" binding:"required"`
	DeptName string `json:"deptName" binding:"required"`
	ParentID uint   `json:"parentId"`
	Path     string `json:"path"`
	Sort     uint   `json:"sort"`
	Status   bool   `json:"status"`
}

// DeptResp 部门响应
type DeptResp struct {
	DeptModel
	Children []*DeptResp `json:"children,omitempty"`
}

// DeptTreeResp 部门树响应
type DeptTreeResp struct {
	ID       uint           `json:"id"`
	DeptName string         `json:"deptName"`
	ParentID uint           `json:"parentId"`
	Path     string         `json:"path"`
	Sort     uint           `json:"sort"`
	Status   bool           `json:"status"`
	Children []*DeptTreeResp `json:"children,omitempty"`
}

// DeptListReq 部门列表请求
type DeptListReq struct {
	DeptName string `json:"deptName" form:"deptName"`
	Status   *bool  `json:"status" form:"status"`
}
