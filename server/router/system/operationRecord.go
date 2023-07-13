package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type OperationRecordRouter struct{}

func (o *OperationRecordRouter) InitOperationRecordRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	operationRecordRouter := Router.Group("or")
	operationRecordApi := api.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("getOrList", operationRecordApi.GetOperationRecordList)
		operationRecordRouter.POST("deleteOr", operationRecordApi.DeleteOperationRecord)
		operationRecordRouter.POST("deleteOrByIds", operationRecordApi.DeleteOperationRecordByIds)
	}
	return operationRecordRouter
}
