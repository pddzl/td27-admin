package monitor

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type OperationLogRouter struct{}

func (o *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	operationLogRouter := Router.Group("opl")
	operationLogApi := api.ApiGroupApp.Monitor.OperationLogApi
	{
		operationLogRouter.POST("getOplList", operationLogApi.GetOperationLogList)
		operationLogRouter.POST("deleteOpl", operationLogApi.DeleteOperationLog)
		operationLogRouter.POST("deleteOplByIds", operationLogApi.DeleteOperationLogByIds)
	}
	return operationLogRouter
}
