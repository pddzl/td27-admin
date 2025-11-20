package request

import (
	"server/internal/model/common/request"
)

type DictDetailSearchParams struct {
	request.PageInfo
	DictID uint `json:"dictID" binding:"required"`
}

type DictDetailFlatReq struct {
	DictID uint `json:"dictID" binding:"required"`
}
