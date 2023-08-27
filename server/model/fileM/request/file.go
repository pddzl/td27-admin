package request

import "server/model/common/request"

// FileSearchParams file分页条件查询
type FileSearchParams struct {
	request.PageInfo
	Name     string `json:"name"`
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
