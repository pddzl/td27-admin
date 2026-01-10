package request

import (
	"server/internal/model/common"
)

// FileSearchParams file分页条件查询
type FileSearchParams struct {
	common.PageInfo
	Name     string `json:"name"`
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
