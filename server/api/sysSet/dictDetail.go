package sysSet

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
	modelSysSet "server/model/sysSet"
	sysSetReq "server/model/sysSet/request"
)

type DictDetailApi struct{}

func (dda *DictDetailApi) GetDictDetail(c *gin.Context) {
	var ddsParams sysSetReq.DictDetailSearchParams
	if err := c.ShouldBindJSON(&ddsParams); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if list, total, err := dictDetailService.GetDictDetail(ddsParams); err != nil {
		commonRes.FailWithMessage("failed", c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
			List:     list,
			Total:    total,
			Page:     ddsParams.Page,
			PageSize: ddsParams.PageSize,
		}, "success", c)
	}
}

func (dda *DictDetailApi) AddDictDetail(c *gin.Context) {
	var dictDetailModel modelSysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if instance, err := dictDetailService.AddDictDetail(&dictDetailModel); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("add failed", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(instance, "success", c)
	}
}

func (dda *DictDetailApi) DelDictDetail(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := dictDetailService.DelDictDetail(cId.ID); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("delete failed", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

func (dda *DictDetailApi) EditDictDetail(c *gin.Context) {
	var dictDetailModel modelSysSet.DictDetailModel
	if err := c.ShouldBindJSON(&dictDetailModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if instance, err := dictDetailService.EditDictDetail(&dictDetailModel); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("edit failed", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(instance, "success", c)
	}
}
