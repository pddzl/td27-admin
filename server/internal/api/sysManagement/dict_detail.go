package sysManagement

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type DictDetailApi struct {
	dictDetailService *serviceSysManagement.DictDetailService
}

func NewDictDetailApi() *DictDetailApi {
	return &DictDetailApi{
		dictDetailService: serviceSysManagement.NewDictDetailService(),
	}
}

func (a *DictDetailApi) List(c *gin.Context) {
	var req modelSysManagement.ListDictDetailReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := a.dictDetailService.List(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("get failed", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "success", c)
	}
}

func (a *DictDetailApi) Flat(c *gin.Context) {
	var req modelSysManagement.DictDetailFlatReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, err := a.dictDetailService.Flat(req.DictID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("flat failed", zap.Error(err))
	} else {
		common.OkWithDetailed(list, "success", c)
	}
}

func (a *DictDetailApi) Create(c *gin.Context) {
	var req modelSysManagement.DictDetailModel
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := a.dictDetailService.Create(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("create failed", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "success", c)
	}
}

func (a *DictDetailApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.dictDetailService.Delete(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("delete failed", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

func (a *DictDetailApi) Update(c *gin.Context) {
	var req modelSysManagement.DictDetailModel
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := a.dictDetailService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("update failed", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "success", c)
	}
}
