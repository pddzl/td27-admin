package sysMonitor

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"server/internal/model/common"
	serviceMonitor "server/internal/service/sysMonitor"
)

type DashboardApi struct {
	dashboardService *serviceMonitor.DashboardService
}

func NewDashboardApi() *DashboardApi {
	return &DashboardApi{
		dashboardService: serviceMonitor.NewDashboardService(),
	}
}

// GetStatistics
// @Tags      DashboardApi
// @Summary   获取仪表盘统计数据
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  common.Response{data=serviceMonitor.DashboardStats,msg=string}
// @Router    /dashboard/statistics [get]
func (a *DashboardApi) GetStatistics(c *gin.Context) {
	stats, err := a.dashboardService.GetStatistics()
	if err != nil {
		slog.Error("获取仪表盘统计数据失败", "error", err)
		common.FailWithMessage("获取统计数据失败", c)
		return
	}
	common.OkWithData(stats, c)
}

// GetRecentOperations
// @Tags      DashboardApi
// @Summary   获取最近操作记录
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  common.Response{data=[]serviceMonitor.RecentOperation,msg=string}
// @Router    /dashboard/recent-operations [get]
func (a *DashboardApi) GetRecentOperations(c *gin.Context) {
	operations, err := a.dashboardService.GetRecentOperations(10)
	if err != nil {
		slog.Error("获取最近操作记录失败", "error", err)
		common.FailWithMessage("获取最近操作记录失败", c)
		return
	}
	common.OkWithData(operations, c)
}

// GetSystemInfo
// @Tags      DashboardApi
// @Summary   获取系统信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  common.Response{data=serviceMonitor.SystemInfo,msg=string}
// @Router    /dashboard/system-info [get]
func (a *DashboardApi) GetSystemInfo(c *gin.Context) {
	info := a.dashboardService.GetSystemInfo()
	common.OkWithData(info, c)
}
