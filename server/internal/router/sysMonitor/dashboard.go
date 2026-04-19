package sysMonitor

import (
	"github.com/gin-gonic/gin"
	"server/internal/middleware"

	apiMonitor "server/internal/api/sysMonitor"
)

type DashboardRouter struct {
	dashboardApi *apiMonitor.DashboardApi
}

func NewDashboardRouter() *DashboardRouter {
	return &DashboardRouter{
		dashboardApi: apiMonitor.NewDashboardApi(),
	}
}

func (r *DashboardRouter) InitDashboardRouter(rg *gin.RouterGroup) {
	base := rg.Group("dashboard")
	record := base.Use(middleware.OperationRecord())
	// record
	record.GET("statistics", r.dashboardApi.GetStatistics)
	record.GET("recent-operations", r.dashboardApi.GetRecentOperations)
	record.GET("system-info", r.dashboardApi.GetSystemInfo)
}
