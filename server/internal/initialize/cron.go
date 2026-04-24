package initialize

import (

	"server/internal/global"
	cronService "server/internal/service/sysTool/cron"
	"log/slog"
)

// InitCron 初始化定时任务调度器
func InitCron() {
	scheduler := cronService.NewScheduler()
	scheduler.Start()
	global.TD27_CRON = scheduler

	// 从数据库加载已启用的任务
	if err := scheduler.LoadFromDB(); err != nil {
		slog.Error("[CRON] load from db failed", "error", err)
	}
}
