package sysManagement

import "server/internal/model/common"

// ApiTreeResp API树响应
type ApiTreeResp struct {
	List       []*ApiTreeNode `json:"list"`
	CheckedIds []uint         `json:"checkedIds"`
}

// UpdateApiReq 更新API请求
type UpdateApiReq struct {
	ID uint `json:"id" binding:"required"`
	CreateApiReq
}

// ListApiReq API列表请求
type ListApiReq struct {
	common.PageInfo
	Path    string `json:"path" form:"path"`
	Method  string `json:"method" form:"method"`
	GroupEn string `json:"group_en" form:"group_en"`
}

// CreateApiReq 创建API请求
type CreateApiReq struct {
	Path        string `json:"path" binding:"required"`
	Method      string `json:"method" binding:"required,oneof=GET POST PUT PATCH DELETE"`
	GroupEN     string `json:"group_en" binding:"required"`
	GroupCN     string `json:"group_cn" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// ApiTreeNode API树节点
type ApiTreeNode struct {
	Id          uint           `json:"id"`
	Key         string         `json:"key"`
	Path        string         `json:"path,omitempty"`
	Method      string         `json:"method,omitempty" `
	GroupEN     string         `json:"group_en,omitempty" `
	GroupCN     string         `json:"group_cn,omitempty" `
	Description string         `json:"description,omitempty"`
	Children    []*ApiTreeNode `json:"children,omitempty"`
}
