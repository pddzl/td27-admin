package sysMonitor

import (
	"runtime"
	"time"

	"server/internal/global"
	modelManagement "server/internal/model/sysManagement"
	modelMonitor "server/internal/model/sysMonitor"
	modelTool "server/internal/model/sysTool"
)

type DashboardService struct{}

func NewDashboardService() *DashboardService {
	return &DashboardService{}
}

// DashboardStats 仪表盘统计数据
type DashboardStats struct {
	UserCount       int64 `json:"userCount"`       // 用户总数
	RoleCount       int64 `json:"roleCount"`       // 角色总数
	DeptCount       int64 `json:"deptCount"`       // 部门总数
	OperationCount  int64 `json:"operationCount"`  // 今日操作数
	CronCount       int64 `json:"cronCount"`       // 定时任务数
	ActiveCronCount int64 `json:"activeCronCount"` // 活跃定时任务数
}

// RecentOperation 最近操作记录
type RecentOperation struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"userName"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Status    int       `json:"status"`
	RespTime  int64     `json:"respTime"`
	CreatedAt time.Time `json:"createdAt"`
}

// SystemInfo 系统信息
type SystemInfo struct {
	AppName     string `json:"appName"`
	Version     string `json:"version"`
	GoVersion   string `json:"goVersion"`
	OS          string `json:"os"`
	Arch        string `json:"arch"`
	NumCPU      int    `json:"numCpu"`
	NumGoroutine int   `json:"numGoroutine"`
	StartTime   string `json:"startTime"`
}

// GetStatistics 获取统计数据
func (s *DashboardService) GetStatistics() (*DashboardStats, error) {
	stats := &DashboardStats{}
	
	// 用户总数
	if err := global.TD27_DB.Model(&modelManagement.UserModel{}).Count(&stats.UserCount).Error; err != nil {
		return nil, err
	}
	
	// 角色总数
	if err := global.TD27_DB.Model(&modelManagement.RoleModel{}).Count(&stats.RoleCount).Error; err != nil {
		return nil, err
	}
	
	// 部门总数
	if err := global.TD27_DB.Model(&modelManagement.DeptModel{}).Count(&stats.DeptCount).Error; err != nil {
		return nil, err
	}
	
	// 今日操作数
	today := time.Now().Format("2006-01-02")
	if err := global.TD27_DB.Model(&modelMonitor.OperationLogModel{}).
		Where("DATE(created_at) = ?", today).
		Count(&stats.OperationCount).Error; err != nil {
		return nil, err
	}
	
	// 定时任务数
	if err := global.TD27_DB.Model(&modelTool.CronModel{}).Count(&stats.CronCount).Error; err != nil {
		return nil, err
	}
	
	// 活跃定时任务数
	if err := global.TD27_DB.Model(&modelTool.CronModel{}).
		Where("open = ?", true).
		Count(&stats.ActiveCronCount).Error; err != nil {
		return nil, err
	}
	
	return stats, nil
}

// GetRecentOperations 获取最近操作记录
func (s *DashboardService) GetRecentOperations(limit int) ([]RecentOperation, error) {
	var logs []modelMonitor.OperationLogModel
	if err := global.TD27_DB.
		Select("id", "user_name", "path", "method", "status", "resp_time", "created_at").
		Order("created_at DESC").
		Limit(limit).
		Find(&logs).Error; err != nil {
		return nil, err
	}
	
	operations := make([]RecentOperation, len(logs))
	for i, log := range logs {
		operations[i] = RecentOperation{
			ID:        log.ID,
			UserName:  log.UserName,
			Path:      log.Path,
			Method:    log.Method,
			Status:    log.Status,
			RespTime:  log.RespTime,
			CreatedAt: log.CreatedAt,
		}
	}
	
	return operations, nil
}

// GetSystemInfo 获取系统信息
func (s *DashboardService) GetSystemInfo() *SystemInfo {
	return &SystemInfo{
		AppName:      "TD27 Admin",
		Version:      "v3.0.0",
		GoVersion:    runtime.Version(),
		OS:           runtime.GOOS,
		Arch:         runtime.GOARCH,
		NumCPU:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
		StartTime:    time.Now().Format("2006-01-02 15:04:05"),
	}
}
