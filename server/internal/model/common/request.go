package common

// PageInfo 分页
type PageInfo struct {
	Page     int `json:"page"`     // 页码
	PageSize int `json:"pageSize"` // 每页大小
}

func (p *PageInfo) Normalize() {
	// pagination safety
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
}

func (p *PageInfo) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// CId 主键ID
type CId struct {
	ID uint `json:"id" binding:"required"` // 主键 ID
}

type CIds struct {
	IDs []uint `json:"ids" binding:"required"`
}
