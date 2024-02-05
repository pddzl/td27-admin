package monitor

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type OperationLogRouter struct{}

func (o *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	operationLogRouter := Router.Group("opl").Use(middleware.OperationRecord())
	operationLogWithoutRouter := Router.Group("opl")

	operationLogApi := api.ApiGroupApp.Monitor.OperationLogApi
	{
		operationLogRouter.POST("deleteOpl", operationLogApi.DeleteOperationLog)
		operationLogRouter.POST("deleteOplByIds", operationLogApi.DeleteOperationLogByIds)
	}
	{
		operationLogWithoutRouter.POST("getOplList", operationLogApi.GetOperationLogList)
	}
}
