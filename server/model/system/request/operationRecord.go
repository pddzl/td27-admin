package request

import "server/model/common/request"

type orStruct struct {
	Method string // 请求方法
	Path   string // 请求路径
	Status int    // http code
}

// OrSearchParams api分页条件查询及排序结构体
type OrSearchParams struct {
	orStruct
	request.PageInfo
	Asc bool `json:"asc"` // 排序方式:升序true|降序true(默认)
}
