package rbac

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"server/internal/global"
	"server/internal/pkg/cache"
	"log/slog"
)

const (
	// PermissionCachePrefix 权限缓存前缀
	PermissionCachePrefix = "rbac:perm:"
	// RolePermissionCachePrefix 角色权限缓存前缀
	RolePermissionCachePrefix = "rbac:role:perm:"
	// UserPermissionCachePrefix 用户权限缓存前缀
	UserPermissionCachePrefix = "rbac:user:perm:"
	// PermissionCacheDuration 权限缓存时间
	PermissionCacheDuration = 30 * time.Minute
)

// PermissionCache 权限缓存
type PermissionCache struct {
	cache *cache.PGCache
}

// NewPermissionCache 创建权限缓存
func NewPermissionCache() *PermissionCache {
	return &PermissionCache{
		cache: cache.NewPGCache(),
	}
}

// CacheUserPermissions 缓存用户权限
func (pc *PermissionCache) CacheUserPermissions(userID uint, username string, permissions []string) error {
	key := fmt.Sprintf("%s%d", UserPermissionCachePrefix, userID)
	data, err := json.Marshal(permissions)
	if err != nil {
		return err
	}
	return pc.cache.Set(context.Background(), key, username, string(data), PermissionCacheDuration)
}

// GetUserPermissions 获取缓存的用户权限
func (pc *PermissionCache) GetUserPermissions(userID uint) ([]string, error) {
	key := fmt.Sprintf("%s%d", UserPermissionCachePrefix, userID)
	data, err := pc.cache.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	var permissions []string
	if err = json.Unmarshal([]byte(data), &permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}

// ClearUserPermissions 清除用户权限缓存
func (pc *PermissionCache) ClearUserPermissions(userID uint) error {
	key := fmt.Sprintf("%s%d", UserPermissionCachePrefix, userID)
	return pc.cache.Del(context.Background(), key)
}

// CacheRolePermissions 缓存角色权限
func (pc *PermissionCache) CacheRolePermissions(roleID uint, username string, permissions []string) error {
	key := fmt.Sprintf("%s%d", RolePermissionCachePrefix, roleID)
	data, err := json.Marshal(permissions)
	if err != nil {
		return err
	}
	return pc.cache.Set(context.Background(), key, username, string(data), PermissionCacheDuration)
}

// GetRolePermissions 获取缓存的角色权限
func (pc *PermissionCache) GetRolePermissions(roleID uint) ([]string, error) {
	key := fmt.Sprintf("%s%d", RolePermissionCachePrefix, roleID)
	data, err := pc.cache.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}

	var permissions []string
	if err := json.Unmarshal([]byte(data), &permissions); err != nil {
		return nil, err
	}
	return permissions, nil
}

// ClearRolePermissions 清除角色权限缓存
func (pc *PermissionCache) ClearRolePermissions(roleID uint) error {
	key := fmt.Sprintf("%s%d", RolePermissionCachePrefix, roleID)
	return pc.cache.Del(context.Background(), key)
}

// ClearAllPermissions 清除所有权限缓存
func (pc *PermissionCache) ClearAllPermissions() error {
	// 由于PostgreSQL缓存没有通配符删除，这里记录日志
	slog.Info("Clear all permission cache - manual cleanup may be required")
	return nil
}

// PermissionCheckResult 权限检查结果
type PermissionCheckResult struct {
	Allowed  bool   `json:"allowed"`
	CacheHit bool   `json:"cacheHit"`
	Reason   string `json:"reason,omitempty"`
}

// CheckPermissionWithCache 带缓存的权限检查
func (pc *PermissionCache) CheckPermissionWithCache(userID uint, resource string, action string) *PermissionCheckResult {
	// 构建权限标识
	permKey := fmt.Sprintf("%s:%s", resource, action)

	// 尝试从缓存获取
	cachedPerms, err := pc.GetUserPermissions(userID)
	if err == nil {
		// 缓存命中
		for _, perm := range cachedPerms {
			if perm == permKey || perm == resource+":*" {
				return &PermissionCheckResult{
					Allowed:  true,
					CacheHit: true,
				}
			}
		}
		return &PermissionCheckResult{
			Allowed:  false,
			CacheHit: true,
			Reason:   "permission not found in cache",
		}
	}

	// 缓存未命中，返回nil让调用方查询数据库
	return &PermissionCheckResult{
		Allowed:  false,
		CacheHit: false,
		Reason:   "cache miss",
	}
}
