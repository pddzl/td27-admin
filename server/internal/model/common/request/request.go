package request

// PageInfo 分页
type PageInfo struct {
	Page     int `json:"page"`     // 页码
	PageSize int `json:"pageSize"` // 每页大小
}

// CId 主键ID
type CId struct {
	ID uint `json:"id" binding:"required"` // 主键ID
}

type CIds struct {
	IDs []uint `json:"ids" binding:"required"`
}
