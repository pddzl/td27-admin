package sysSet

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common/request"
	"server/internal/model/common/response"
	modelSysSet "server/internal/model/entity/sysSet"
	"server/internal/service/sysSet"
)

type DictApi struct {
	dictService *sysSet.DictService
}

func NewDictApi() *DictApi {
	return &DictApi{
		dictService: sysSet.NewDictService(),
	}
}

func (da *DictApi) GetDict(c *gin.Context) {
	if list, err := da.dictService.GetDict(); err != nil {
		response.Fail(c)
		global.TD27_LOG.Error("Get Error", zap.Error(err))
	} else {
		response.OkWithDetailed(list, "success", c)
	}
}

func (da *DictApi) AddDict(c *gin.Context) {
	var dictModel modelSysSet.DictModel
	if err := c.ShouldBindJSON(&dictModel); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if instance, err := da.dictService.AddDict(&dictModel); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("failed", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "success", c)
	}
}

func (da *DictApi) DelDict(c *gin.Context) {
	var cId request.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if err := da.dictService.DelDict(cId.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("Delete error", zap.Error(err))
	} else {
		response.Ok(c)
	}
}

func (da *DictApi) EditDict(c *gin.Context) {
	var dictModel modelSysSet.DictModel
	if err := c.ShouldBindJSON(&dictModel); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if instance, err := da.dictService.EditDict(&dictModel); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("edit failed", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "success", c)
	}
}
