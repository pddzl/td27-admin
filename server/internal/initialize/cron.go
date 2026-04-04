package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"server/internal/global"
	modelSysTool "server/internal/model/sysTool"
)

// InitCron 初始化Cron
func InitCron() *cron.Cron {
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
