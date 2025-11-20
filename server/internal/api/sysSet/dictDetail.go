package sysSet

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common/request"
	"server/internal/model/common/response"
	modelSysSet "server/internal/model/entity/sysSet"
	sysSetReq "server/internal/model/entity/sysSet/request"
	"server/internal/service/sysSet"
)

type DictDetailApi struct {
	dictDetailService *sysSet.DictDetailService
}

func NewDictDetailApi() *DictDetailApi {
	return &DictDetailApi{
		dictDetailService: sysSet.NewDictDetailService(),
	}
}

func (dda *DictDetailApi) GetDictDetail(c *gin.Context) {
	var ddsParams sysSetReq.DictDetailSearchParams
	if err := c.ShouldBindJSON(&ddsParams); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if list, total, err := dda.dictDetailService.GetDictDetail(ddsParams); err != nil {
		response.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		response.OkWithDetailed(response.Page{
			List:     list,
			Total:    total,
			Page:     ddsParams.Page,
			PageSize: ddsParams.PageSize,
		}, "success", c)
	}
}

func (dda *DictDetailApi) GetDictDetailFlat(c *gin.Context) {
	var flatReq sysSetReq.DictDetailFlatReq
	if err := c.ShouldBindJSON(&flatReq); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if list, err := dda.dictDetailService.GetDictDetailFlat(flatReq.DictID); err != nil {
		response.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		response.OkWithDetailed(list, "success", c)
	}
}

func (dda *DictDetailApi) AddDictDetail(c *gin.Context) {
	var dictDetailModel modelSysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if instance, err := dda.dictDetailService.AddDictDetail(&dictDetailModel); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("add failed", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "success", c)
	}
}

func (dda *DictDetailApi) DelDictDetail(c *gin.Context) {
	var cId request.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if err := dda.dictDetailService.DelDictDetail(cId.ID); err != nil {
		response.Fail(c)
		global.TD27_LOG.Error("delete failed", zap.Error(err))
	} else {
		response.Ok(c)
	}
}

func (dda *DictDetailApi) EditDictDetail(c *gin.Context) {
	var dictDetailModel modelSysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if instance, err := dda.dictDetailService.EditDictDetail(&dictDetailModel); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("edit failed", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "success", c)
	}
}
