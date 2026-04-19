package sysMonitor

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysMonitor"
)

type OperationLogRouter struct {
	operationLogApi *sysMonitor.OperationLogApi
}

func NewOperationLogRouter() *OperationLogRouter {
	return &OperationLogRouter{operationLogApi: sysMonitor.NewOperationLogApi()}
}

func (r *OperationLogRouter) InitOperationLogRouter(rg *gin.RouterGroup) {
	base := rg.Group("opl")
	base.POST("delete", r.operationLogApi.Delete)
	base.POST("deleteByIds", r.operationLogApi.DeleteByIds)
	base.POST("list", r.operationLogApi.List)
}
