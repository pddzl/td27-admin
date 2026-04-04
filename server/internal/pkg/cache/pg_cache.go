package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"server/internal/global"
	"server/internal/model/sysTool"
)

// PGCache PostgreSQL 缓存实现
type PGCache struct{}

// NewPGCache 创建 PostgreSQL 缓存实例
func NewPGCache() *PGCache {
	return &PGCache{}
}

// getDB 获取数据库连接（延迟获取，避免初始化时 nil）
func (c *PGCache) getDB() *gorm.DB {
	return global.TD27_DB
}

// Get 获取缓存值
func (c *PGCache) Get(ctx context.Context, key string) (string, error) {
	db := c.getDB()
	if db == nil {
		return "", fmt.Errorf("database not initialized")
	}

	var cache sysTool.CacheModel
	err := db.WithContext(ctx).
		Where("\"key\" = ? AND expires_at > ?", key, time.Now()).
		First(&cache).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("cache miss")
		}
		return "", err
	}

	return cache.Value, nil
}

// Set 设置缓存值
func (c *PGCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	db := c.getDB()
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	expiresAt := time.Now().Add(expiration)

	// 先尝试更新
	result := db.WithContext(ctx).Exec(`
		UPDATE sys_tool_cache 
		SET "value" = ?, expires_at = ?, updated_at = ?
		WHERE "key" = ? and "deleted_at" = null
	`, value, expiresAt, time.Now(), key)

	// 如果没有记录被更新，则插入
	if result.RowsAffected == 0 {
		return db.WithContext(ctx).Exec(`
			INSERT INTO sys_tool_cache ("key", "value", expires_at, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)
		`, key, value, expiresAt, time.Now(), time.Now()).Error
	}

	return result.Error
}

// Del 删除缓存
func (c *PGCache) Del(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}

	db := c.getDB()
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	return db.WithContext(ctx).
		Where("\"key\" IN ?", keys).Unscoped().
		Delete(&sysTool.CacheModel{}).Error
}

// CleanupExpired 清理过期缓存（可由定时任务调用）
func (c *PGCache) CleanupExpired(ctx context.Context) error {
	db := c.getDB()
	if db == nil {
		return fmt.Errorf("database not initialized")
	}

	result := db.WithContext(ctx).
		Where("expires_at <= ?", time.Now()).
		Delete(&sysTool.CacheModel{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		global.TD27_LOG.Info("清理过期缓存", zap.Int64("count", result.RowsAffected))
	}

	return nil
}

// Exists 检查 key 是否存在
func (c *PGCache) Exists(ctx context.Context, key string) bool {
	db := c.getDB()
	if db == nil {
		return false
	}

	var count int64
	db.WithContext(ctx).
		Model(&sysTool.CacheModel{}).
		Where("\"key\" = ? AND expires_at > ?", key, time.Now()).
		Count(&count)
	return count > 0
}

// TTL 获取 key 的剩余过期时间
func (c *PGCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	db := c.getDB()
	if db == nil {
		return 0, fmt.Errorf("database not initialized")
	}

	var cache sysTool.CacheModel
	err := db.WithContext(ctx).
		Where("\"key\" = ? AND expires_at > ?", key, time.Now()).
		Select("expires_at").
		First(&cache).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return -1, fmt.Errorf("key not found or expired")
		}
		return 0, err
	}

	return time.Until(cache.ExpiresAt), nil
}

// ListKeysByPrefix 根据前缀获取所有key（用于多登录模式查找用户的所有token）
func (c *PGCache) ListKeysByPrefix(ctx context.Context, prefix string) ([]string, error) {
	db := c.getDB()
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var caches []sysTool.CacheModel
	err := db.WithContext(ctx).
		Where("\"key\" LIKE ? AND expires_at > ?", prefix+"%", time.Now()).
		Select("\"key\"").
		Find(&caches).Error

	if err != nil {
		return nil, err
	}

	keys := make([]string, len(caches))
	for i, cache := range caches {
		keys[i] = cache.Key
	}

	return keys, nil
}
