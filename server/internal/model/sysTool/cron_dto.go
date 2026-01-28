package sysTool

type SwitchReq struct {
	ID   uint `json:"id" binding:"required"`
	Open bool `json:"open"`
}
