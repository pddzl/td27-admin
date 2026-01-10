package sysSet

import (
	"server/internal/model/common"
	entitySysSet "server/internal/model/sysSet"
	sysSetReq "server/internal/model/sysSet/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	serviceSysSet "server/internal/service/sysSet"
)

type DictDetailApi struct {
	dictDetailService *serviceSysSet.DictDetailService
}

func NewDictDetailApi() *DictDetailApi {
	return &DictDetailApi{
		dictDetailService: serviceSysSet.NewDictDetailService(),
	}
}

func (dda *DictDetailApi) GetDictDetail(c *gin.Context) {
	var ddsParams sysSetReq.DictDetailSearchParams
	if err := c.ShouldBindJSON(&ddsParams); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := dda.dictDetailService.GetDictDetail(ddsParams); err != nil {
		common.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
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
		common.FailReq(err.Error(), c)
		return
	}

	if list, err := dda.dictDetailService.GetDictDetailFlat(flatReq.DictID); err != nil {
		common.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		common.OkWithDetailed(list, "success", c)
	}
}

func (dda *DictDetailApi) AddDictDetail(c *gin.Context) {
	var dictDetailModel entitySysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := dda.dictDetailService.AddDictDetail(&dictDetailModel); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("add failed", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "success", c)
	}
}

func (dda *DictDetailApi) DelDictDetail(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := dda.dictDetailService.DelDictDetail(cId.ID); err != nil {
		common.Fail(c)
		global.TD27_LOG.Error("delete failed", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

func (dda *DictDetailApi) EditDictDetail(c *gin.Context) {
	var dictDetailModel entitySysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := dda.dictDetailService.EditDictDetail(&dictDetailModel); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("edit failed", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "success", c)
	}
}
