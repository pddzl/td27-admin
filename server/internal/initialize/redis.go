package initialize

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"server/internal/global"
)

func Redis() *redis.Client {
	redisCfg := global.TD27_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.TD27_CONFIG.Redis.Host, global.TD27_CONFIG.Redis.Port),
		//Password: redisCfg.Password, // no password set
		DB: redisCfg.DB, // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		global.TD27_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.TD27_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}

	return client
}
