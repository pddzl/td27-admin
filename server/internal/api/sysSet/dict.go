package sysSet

import (
	commonReq "server/internal/model/common"
	entitySysSet "server/internal/model/sysSet"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	serviceSysyset "server/internal/service/sysSet"
)

type DictApi struct {
	dictService *serviceSysyset.DictService
}

func NewDictApi() *DictApi {
	return &DictApi{
		dictService: serviceSysyset.NewDictService(),
	}
}

func (da *DictApi) GetDict(c *gin.Context) {
	if list, err := da.dictService.GetDict(); err != nil {
		commonReq.Fail(c)
		global.TD27_LOG.Error("Get Error", zap.Error(err))
	} else {
		commonReq.OkWithDetailed(list, "success", c)
	}
}

func (da *DictApi) AddDict(c *gin.Context) {
	var dictModel entitySysSet.DictModel
	if err := c.ShouldBindJSON(&dictModel); err != nil {
		commonReq.FailReq(err.Error(), c)
		return
	}

	if instance, err := da.dictService.AddDict(&dictModel); err != nil {
		commonReq.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("failed", zap.Error(err))
	} else {
		commonReq.OkWithDetailed(instance, "success", c)
	}
}

func (da *DictApi) DelDict(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonReq.FailReq(err.Error(), c)
		return
	}

	if err := da.dictService.DelDict(cId.ID); err != nil {
		commonReq.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Delete error", zap.Error(err))
	} else {
		commonReq.Ok(c)
	}
}

func (da *DictApi) EditDict(c *gin.Context) {
	var dictModel entitySysSet.DictModel
	if err := c.ShouldBindJSON(&dictModel); err != nil {
		commonReq.FailReq(err.Error(), c)
		return
	}

	if instance, err := da.dictService.EditDict(&dictModel); err != nil {
		commonReq.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("edit failed", zap.Error(err))
	} else {
		commonReq.OkWithDetailed(instance, "success", c)
	}
}
