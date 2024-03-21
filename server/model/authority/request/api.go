package request

import "server/model/common/request"

type apiStruct struct {
	Path        string `json:"path"`                                                 // 路径
	ApiGroup    string `json:"api_group"`                                            // API分组
	Method      string `json:"method" binding:"omitempty,oneof=GET POST DELETE PUT"` // 请求方法
	Description string `json:"description"`                                          // 描述
}

// ApiSearchParams api分页条件查询及排序结构体
type ApiSearchParams struct {
	apiStruct
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
