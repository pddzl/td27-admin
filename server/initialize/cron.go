package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"server/config"
	"server/global"
	modelSysTool "server/model/sysTool"
	"server/utils"
)

// Crontab 添加计划任务
func crontab() {
	if global.TD27_CONFIG.Crontab.Open {
		ct := cron.New()
		for index := range global.TD27_CONFIG.Crontab.Objects {
			go func(cObject config.Object) {
				_, err := ct.AddFunc(global.TD27_CONFIG.Crontab.Spec, func() {
					err := utils.ClearTable(global.TD27_DB, cObject.TableName, cObject.CompareField, cObject.Interval)
					if err != nil {
						global.TD27_LOG.Error("clear table", zap.Error(err))
					}
				})
				if err != nil {
					global.TD27_LOG.Error("cron add func", zap.Error(err))
				}
			}(global.TD27_CONFIG.Crontab.Objects[index])
		}
		// 启动cron
		ct.Start()
	}
}

// InitCron 初始化Cron
func InitCron() *cron.Cron {
	// 配置文件方式cron
	crontab()
	// 页面方式配置
	instance := cron.New(cron.WithSeconds()) // 支持秒
	instance.Start()                         // 启动cron
	return instance
}

func CheckCron() {
	var cronModelList []modelSysTool.CronModel
	global.TD27_DB.Where("open = ?", 1).Find(&cronModelList)
	for _, cronModel := range cronModelList {
		entryId, err := global.TD27_CRON.AddJob(cronModel.Expression, &cronModel)
		if err != nil {
			global.TD27_LOG.Error("CRON", zap.Error(err))
		} else {
			global.TD27_DB.Model(cronModel).Update("entryId", entryId)
		}
	}
}
