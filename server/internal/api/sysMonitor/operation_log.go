package sysMonitor

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelMonitor "server/internal/model/sysMonitor"
	serviceMonitor "server/internal/service/sysMonitor"
)

type OperationLogApi struct {
	operationLogService *serviceMonitor.OperationLogService
}

func NewOperationLogApi() *OperationLogApi {
	return &OperationLogApi{
		operationLogService: serviceMonitor.NewOperationLogService(),
	}
}

// List
// @Tags      OperationLogApi
// @Summary   分页获取操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelMonitor.OrListReq true "请求参数"
// @Success   200   {object}  common.Response{data=common.Page{list=[]modelMonitor.OperationLogModel},msg=string}
// @Router    /opl/list [post]
func (a *OperationLogApi) List(c *gin.Context) {
	var req modelMonitor.OrListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := a.operationLogService.List(&req); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// Delete
// @Tags      OperationLogApi
// @Summary   删除操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /opl/delete [post]
func (a *OperationLogApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.operationLogService.Delete(cId.ID); err != nil {
		common.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds
// @Tags      OperationLogApi
// @Summary   批量删除操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CIds true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /opl/deleteByIds [post]
func (a *OperationLogApi) DeleteByIds(c *gin.Context) {
	var cIds common.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	// todo
	// response delete row nums
	if _, err := a.operationLogService.DeleteByIds(cIds.IDs); err != nil {
		common.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除成功", c)
	}
}
