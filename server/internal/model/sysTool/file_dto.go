package sysTool

import (
	"server/internal/model/common"
)

// ListFileReq file分页条件查询
type ListFileReq struct {
	common.PageInfo
	Name string `json:"name"`
	Desc bool   `json:"desc"` // 排序方式:升序false(默认)|降序true
}
