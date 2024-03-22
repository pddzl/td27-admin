package monitor

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
	monitorReq "server/model/monitor/request"
)

type OperationLogApi struct{}

// GetOperationLogList
// @Tags      OperationLogApi
// @Summary   分页获取操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      monitorReq.OrSearchParams true "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]monitor.OperationLogModel},msg=string}
// @Router    /opl/getOperationLogList [post]
func (o *OperationLogApi) GetOperationLogList(c *gin.Context) {
	var orSp monitorReq.OrSearchParams
	if err := c.ShouldBindJSON(&orSp); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if list, total, err := operationLogService.GetOperationLogList(orSp); err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
			List:     list,
			Total:    total,
			Page:     orSp.Page,
			PageSize: orSp.PageSize,
		}, "获取成功", c)
	}
}

// DeleteOperationLog
// @Tags      OperationLogApi
// @Summary   删除操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /opl/deleteOperationLog [post]
func (o *OperationLogApi) DeleteOperationLog(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := operationLogService.DeleteOperationLog(cId.ID); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}

// DeleteOperationLogByIds
// @Tags      OperationLogApi
// @Summary   批量删除操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CIds true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /opl/deleteOperationLogByIds [post]
func (o *OperationLogApi) DeleteOperationLogByIds(c *gin.Context) {
	var cIds commonReq.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := operationLogService.DeleteOperationLogByIds(cIds.IDs); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}
