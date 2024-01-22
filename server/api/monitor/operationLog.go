package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/global"
	authorityReq "server/model/authority/request"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
)

type OperationLogApi struct{}

// GetOperationLogList
// @Tags      OperationLogApi
// @Summary   分页获取操作记录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.OrSearchParams true "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]monitor.OperationLogModel},msg=string}
// @Router    /opl/getOperationLogList [post]
func (o *OperationLogApi) GetOperationLogList(c *gin.Context) {
	var orSp authorityReq.OrSearchParams
	_ = c.ShouldBindJSON(&orSp)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&orSp); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
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
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
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
	_ = c.ShouldBindJSON(&cIds)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cIds); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := operationLogService.DeleteOperationLogByIds(cIds.Ids); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}
