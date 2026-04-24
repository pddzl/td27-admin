package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"server/configs"
)

var (
	TD27_VP                  *viper.Viper
	TD27_CONFIG              configs.Server
	TD27_LOG                 *zap.Logger
	TD27_DB                  *gorm.DB
	TD27_Concurrency_Control = &singleflight.Group{}
	TD27_CRON                interface{}
	// Scheduler is set by initialize.InitCron() — use cronService.GetScheduler() to access
)
