package request

// CId 主键ID
type CId struct {
	ID uint `json:"id" validate:"required"` // 主键ID
}

type CIds struct {
	Ids []uint `json:"ids" validate:"required"`
}
