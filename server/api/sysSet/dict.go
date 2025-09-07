package sysSet

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
	modelSysSet "server/model/sysSet"
)

type DictApi struct{}

func (da *DictApi) GetDict(c *gin.Context) {
	if list, err := dictService.GetDict(); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("Get Error", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(list, "Get Success", c)
	}
}

func (da *DictApi) AddDict(c *gin.Context) {
	var dictModel modelSysSet.DictModel
	if err := c.ShouldBindJSON(&dictModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if instance, err := dictService.AddDict(&dictModel); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Add Error", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(instance, "Add Success", c)
	}
}

func (da *DictApi) DelDict(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := dictService.DelDict(cId.ID); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Delete error", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

func (da *DictApi) EditDict(c *gin.Context) {
	var dictModel modelSysSet.DictModel
	if err := c.ShouldBindJSON(&dictModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if instance, err := dictService.EditDict(&dictModel); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Edit Error", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(instance, "Edit Success", c)
	}
}
