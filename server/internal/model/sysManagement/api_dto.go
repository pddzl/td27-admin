package sysManagement

import (
	"server/internal/model/common"
)

type apiStruct struct {
	Path        string `json:"path"`                                                 // 路径
	ApiGroup    string `json:"apiGroup"`                                             // API 分组
	Method      string `json:"method" binding:"omitempty,oneof=GET POST DELETE PUT"` // 请求方法
	Description string `json:"description"`                                          // 描述
}

// ListApiReq api分页条件查询及排序结构体
type ListApiReq struct {
	apiStruct
	common.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type ApiTreeResp struct {
	List       interface{} `json:"list"`
	CheckedKey []string    `json:"checkedKey"`
}
