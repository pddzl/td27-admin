package monitor

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/monitor"
	"server/internal/middleware"
)

type OperationLogRouter struct {
	operationLogApi *monitor.OperationLogApi
}

func NewOperationLogRouter() *OperationLogRouter {
	return &OperationLogRouter{operationLogApi: monitor.NewOperationLogApi()}
}

func (or *OperationLogRouter) InitOperationLogRouter(Router *gin.RouterGroup) {
	// record
	operationLogRouter := Router.Group("opl").Use(middleware.OperationRecord())
	operationLogRouter.POST("deleteOpl", or.operationLogApi.DeleteOperationLog)
	operationLogRouter.POST("deleteOplByIds", or.operationLogApi.DeleteOperationLogByIds)
	// not record
	operationLogWithoutRouter := Router.Group("opl")
	operationLogWithoutRouter.POST("getOplList", or.operationLogApi.GetOperationLogList)
}
