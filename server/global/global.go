package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"server/config"
)

var (
	TD27_VP                  *viper.Viper
	TD27_CONFIG              config.Server
	TD27_LOG                 *zap.Logger
	TD27_DB                  *gorm.DB
	TD27_REDIS               *redis.Client
	TD27_Concurrency_Control = &singleflight.Group{}
)
