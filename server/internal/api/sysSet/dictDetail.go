package sysSet

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	commonReq "server/internal/model/common/request"
	commonResp "server/internal/model/common/response"
	entitySysSet "server/internal/model/entity/sysSet"
	sysSetReq "server/internal/model/entity/sysSet/request"
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if list, total, err := dda.dictDetailService.GetDictDetail(ddsParams); err != nil {
		commonResp.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(commonResp.Page{
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if list, err := dda.dictDetailService.GetDictDetailFlat(flatReq.DictID); err != nil {
		commonResp.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(list, "success", c)
	}
}

func (dda *DictDetailApi) AddDictDetail(c *gin.Context) {
	var dictDetailModel entitySysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if instance, err := dda.dictDetailService.AddDictDetail(&dictDetailModel); err != nil {
		commonResp.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("add failed", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(instance, "success", c)
	}
}

func (dda *DictDetailApi) DelDictDetail(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := dda.dictDetailService.DelDictDetail(cId.ID); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("delete failed", zap.Error(err))
	} else {
		commonResp.Ok(c)
	}
}

func (dda *DictDetailApi) EditDictDetail(c *gin.Context) {
	var dictDetailModel entitySysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if instance, err := dda.dictDetailService.EditDictDetail(&dictDetailModel); err != nil {
		commonResp.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("edit failed", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(instance, "success", c)
	}
}
