package sysSet

import (
	"server/internal/model/common"
)

type DictDetailSearchParams struct {
	common.PageInfo
	DictID uint `json:"dictID" binding:"required"`
}

type DictDetailFlatReq struct {
	DictID uint `json:"dictID" binding:"required"`
}
