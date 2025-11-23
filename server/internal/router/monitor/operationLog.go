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

func (or *OperationLogRouter) InitOperationLogRouter(rg *gin.RouterGroup) {
	base := rg.Group("opl")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("deleteOpl", or.operationLogApi.DeleteOperationLog)
	record.POST("deleteOplByIds", or.operationLogApi.DeleteOperationLogByIds)
	// not record
	base.POST("getOplList", or.operationLogApi.GetOperationLogList)
}
