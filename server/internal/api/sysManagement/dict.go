package sysManagement

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type DictApi struct {
	dictService *serviceSysManagement.DictService
}

func NewDictApi() *DictApi {
	return &DictApi{
		dictService: serviceSysManagement.NewDictService(),
	}
}

func (a *DictApi) List(c *gin.Context) {
	var req common.PageInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, count, err := a.dictService.List(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("list error", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			List:     list,
			Total:    count,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "success", c)
	}
}

func (a *DictApi) Create(c *gin.Context) {
	var req modelSysManagement.DictModel
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := a.dictService.Create(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Create error", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "success", c)
	}
}

func (a *DictApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.dictService.Delete(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Delete error", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

func (a *DictApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateDictReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.dictService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("update failed", zap.Error(err))
	} else {
		common.OkWithMessage("update success", c)
	}
}
