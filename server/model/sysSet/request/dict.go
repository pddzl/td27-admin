package request

import "server/model/common/request"

type DictDetailSearchParams struct {
	request.PageInfo
	DictID uint `json:"dictID" binding:"required"`
}
