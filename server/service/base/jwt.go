package base

import (
	"context"
	"time"

	"server/global"
)

type JwtService struct{}

// GetRedisJWT 获取jwt
func (jwtService *JwtService) GetRedisJWT(username string) (redisJWT string, err error) {
	redisJWT, err = global.TD27_REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(username string, jwt string) (err error) {
	// 此处过期时间等于jwt过期时间
	err = global.TD27_REDIS.Set(context.Background(), username, jwt, time.Duration(global.TD27_CONFIG.JWT.ExpiresTime)*time.Second).Err()
	return err
}
