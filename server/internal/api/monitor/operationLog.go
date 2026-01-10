package monitor

import (
	"server/internal/model/common"
	monitorReq "server/internal/model/monitor/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	_ "server/internal/model/entity/monitor"
	serviceMonitor "server/internal/service/monitor"
)

type OperationLogApi struct {
	operationLogService *serviceMonitor.OperationLogService
}

func NewOperationLogApi() *OperationLogApi {
	return &OperationLogApi{
		operationLogService: serviceMonitor.NewOperationLogService(),
	}
}

// GetOperationLogList
// @Tags      OperationLogApi
// @Summary   分页获取操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      monitorReq.OrSearchParams true "请求参数"
// @Success   200   {object}  commonResp.Response{data=commonResp.Page{list=[]monitor.OperationLogModel},msg=string}
// @Router    /opl/getOperationLogList [post]
func (oa *OperationLogApi) GetOperationLogList(c *gin.Context) {
	var orSp monitorReq.OrSearchParams
	if err := c.ShouldBindJSON(&orSp); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := oa.operationLogService.GetOperationLogList(orSp); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
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
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /opl/deleteOperationLog [post]
func (oa *OperationLogApi) DeleteOperationLog(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := oa.operationLogService.DeleteOperationLog(cId.ID); err != nil {
		common.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除成功", c)
	}
}

// DeleteOperationLogByIds
// @Tags      OperationLogApi
// @Summary   批量删除操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CIds true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /opl/deleteOperationLogByIds [post]
func (oa *OperationLogApi) DeleteOperationLogByIds(c *gin.Context) {
	var cIds common.CIds
	if err := c.ShouldBindJSON(&cIds); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := oa.operationLogService.DeleteOperationLogByIds(cIds.IDs); err != nil {
		common.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除成功", c)
	}
}
