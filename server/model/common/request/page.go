package request

// PageInfo 分页
type PageInfo struct {
	Page     int `json:"page"`     // 页码
	PageSize int `json:"pageSize"` // 每页大小
}
