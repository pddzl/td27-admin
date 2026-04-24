package global

import (
	"log/slog"

	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"server/configs"
)

var (
	TD27_VP                  *viper.Viper
	TD27_CONFIG              configs.Server
	TD27_LOG                 *slog.Logger
	TD27_DB                  *gorm.DB
	TD27_Concurrency_Control = &singleflight.Group{}
	TD27_CRON                interface{}
)
