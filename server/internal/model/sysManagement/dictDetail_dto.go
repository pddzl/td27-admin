package sysManagement

import (
	"server/internal/model/common"
)

type ListDictDetailReq struct {
	common.PageInfo
	DictID uint `json:"dictID" binding:"required"`
}

type DictDetailFlatReq struct {
	DictID uint `json:"dictID" binding:"required"`
}
