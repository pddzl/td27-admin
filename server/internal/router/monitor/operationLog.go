package monitor

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/monitor"
	"server/internal/middleware"
)

type OperationLogRouter struct {
	OperationLogApi *monitor.OperationLogApi
}

func NewOperationLogRouter() *OperationLogRouter {
	return &OperationLogRouter{OperationLogApi: monitor.NewOperationLogApi()}
}

func (or *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	// record
	operationLogRouter := Router.Group("opl").Use(middleware.OperationRecord())
	operationLogRouter.POST("deleteOpl", or.OperationLogApi.DeleteOperationLog)
	operationLogRouter.POST("deleteOplByIds", or.OperationLogApi.DeleteOperationLogByIds)
	// not record
	operationLogWithoutRouter := Router.Group("opl")
	operationLogWithoutRouter.POST("getOplList", or.OperationLogApi.GetOperationLogList)
}
