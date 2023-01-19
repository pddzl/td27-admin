package system

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"

	"server/global"
	"server/model/system"
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

// JoinInBlacklist 拉黑jwt
func (jwtService *JwtService) JoinInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.TD27_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	return
}

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.TD27_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}
