package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"server/config"
	"server/global"
	"server/utils"
)

// Crontab 添加计划任务
func Crontab() {
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
	}
}
