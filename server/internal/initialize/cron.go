package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"server/internal/global"
	modelSysTool "server/internal/model/sysTool"
)

// Crontab 添加计划任务
//func crontab() {
//	if global.TD27_CONFIG.Crontab.Open {
//		ct := cron.New()
//		for index := range global.TD27_CONFIG.Crontab.Objects {
//			go func(cObject configs.Object) {
//				_, err := ct.AddFunc(global.TD27_CONFIG.Crontab.Spec, func() {
//					err := pkg.ClearTable(global.TD27_DB, cObject.TableName, cObject.CompareField, cObject.Interval)
//					if err != nil {
//						global.TD27_LOG.Error("clear table", zap.Error(err))
//					}
//				})
//				if err != nil {
//					global.TD27_LOG.Error("cron add func", zap.Error(err))
//				}
//			}(global.TD27_CONFIG.Crontab.Objects[index])
//		}
//
//		// 添加缓存清理任务（每小时执行一次）
//		_, err := ct.AddFunc("@hourly", func() {
//			pgCache := cache.NewPGCache()
//			if err := pgCache.CleanupExpired(context.Background()); err != nil {
//				global.TD27_LOG.Error("cleanup expired cache", zap.Error(err))
//			}
//		})
//		if err != nil {
//			global.TD27_LOG.Error("cron add cache cleanup func", zap.Error(err))
//		}
//
//		// 启动cron
//		ct.Start()
//	}
//}

// InitCron 初始化Cron
func InitCron() *cron.Cron {
	// 配置文件方式cron
	//crontab()
	// 页面方式配置
	instance := cron.New(cron.WithSeconds()) // 支持秒
	instance.Start()                         // 启动cron
	return instance
}

func CheckCron() {
	var cronModelList []modelSysTool.CronModel
	global.TD27_DB.Where("open = ?", true).Find(&cronModelList)
	for _, cronModel := range cronModelList {
		entryId, err := global.TD27_CRON.AddJob(cronModel.Expression, &cronModel)
		if err != nil {
			global.TD27_LOG.Error("CRON", zap.Error(err))
		} else {
			global.TD27_DB.Model(cronModel).Update("entryId", entryId)
		}
	}
}
