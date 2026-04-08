package sysMonitor

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysMonitor"
	"server/internal/middleware"
)

type OperationLogRouter struct {
	operationLogApi *sysMonitor.OperationLogApi
}

func NewOperationLogRouter() *OperationLogRouter {
	return &OperationLogRouter{operationLogApi: sysMonitor.NewOperationLogApi()}
}

func (r *OperationLogRouter) InitOperationLogRouter(rg *gin.RouterGroup) {
	base := rg.Group("opl")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("delete", r.operationLogApi.Delete)
	record.POST("deleteByIds", r.operationLogApi.DeleteByIds)
	// not record
	base.POST("list", r.operationLogApi.List)
}
