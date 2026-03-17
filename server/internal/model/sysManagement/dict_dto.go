package sysManagement

type UpdateDictReq struct {
	ID     uint   `json:"id"`
	CNName string `json:"cn_name" binding:"required"`
}
