package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemReq "server/model/system/request"
)

type OperationRecordApi struct{}

// GetOperationRecordList 分页获取操作记录
func (o *OperationRecordApi) GetOperationRecordList(c *gin.Context) {
	var orSp systemReq.OrSearchParams
	_ = c.ShouldBindJSON(&orSp)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&orSp); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if list, total, err := operationService.GetOperationRecordList(orSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     orSp.Page,
			PageSize: orSp.PageSize,
		}, "获取成功", c)
	}
}

// DeleteOperationRecord 批量删除操作记录
func (o *OperationRecordApi) DeleteOperationRecord(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := operationService.DeleteOperation(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOperationRecordByIds 批量删除操作记录
func (o *OperationRecordApi) DeleteOperationRecordByIds(c *gin.Context) {
	var cIds request.CIds
	_ = c.ShouldBindJSON(&cIds)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cIds); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := operationService.DeleteOperationByIds(cIds.Ids); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
