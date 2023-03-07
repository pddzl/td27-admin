package request

// CId 主键ID
type CId struct {
	ID uint `json:"id" validate:"required"` // 主键ID
}

// PageInfo 分页
type PageInfo struct {
	Page     int `json:"page" validate:"required"`     // 页码
	PageSize int `json:"pageSize" validate:"required"` // 每页大小
}
