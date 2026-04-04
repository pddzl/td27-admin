package sysMonitor

import (
	"github.com/gin-gonic/gin"

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

func (r *DashboardRouter) InitDashboardRouter(group *gin.RouterGroup) {
	dashboardGroup := group.Group("/dashboard")
	{
		dashboardGroup.GET("/statistics", r.dashboardApi.GetStatistics)
		dashboardGroup.GET("/recent-operations", r.dashboardApi.GetRecentOperations)
		dashboardGroup.GET("/system-info", r.dashboardApi.GetSystemInfo)
	}
}
