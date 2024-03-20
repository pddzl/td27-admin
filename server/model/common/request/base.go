package request

// CId 主键ID
type CId struct {
	ID uint `json:"id" binding:"required"` // 主键ID
}

type CIds struct {
	Ids []uint `json:"ids" binding:"required"`
}
