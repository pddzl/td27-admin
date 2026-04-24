package sysManagement

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"server/internal/global"
	"server/internal/model/sysManagement"
	"server/internal/pkg/cache"
)

const (
	// UserCachePrefix 用户缓存key前缀
	UserCachePrefix = "user:"
	// UserTokenPrefix 用户token列表前缀（多设备登录模式）
	UserTokenPrefix = "user_tokens:"
	// UserSingleTokenKey 单设备登录模式下的token key
	UserSingleTokenKey = "user_single_token:"
	// UserCacheDuration 用户缓存时间（比JWT过期时间稍长）
	UserCacheDuration = 24 * time.Hour
)

type JwtService struct {
	cache *cache.PGCache
}

func NewJwtService() *JwtService {
	return &JwtService{
		cache: cache.NewPGCache(),
	}
}

// isMultiLogin 检查是否允许多设备登录
func (jwtService *JwtService) isMultiLogin() bool {
	return global.TD27_CONFIG.JWT.MultiLogin
}

// getMultiLoginLimit 获取多设备登录数量限制
func (jwtService *JwtService) getMultiLoginLimit() int {
	limit := global.TD27_CONFIG.JWT.MultiLoginLimit
	if limit < 0 {
		return 0
	}
	return limit
}

// generateTokenKey 生成token缓存key
// 单设备模式: user_single_token:{username}
// 多设备模式: user_tokens:{username}:{tokenID}
func (jwtService *JwtService) generateTokenKey(username string, tokenID string) string {
	if jwtService.isMultiLogin() {
		return fmt.Sprintf("%s%s:%s", UserTokenPrefix, username, tokenID)
	}
	return fmt.Sprintf("%s%s", UserSingleTokenKey, username)
}

// getUserTokensPrefix 获取用户token key前缀（用于查找所有token）
func (jwtService *JwtService) getUserTokensPrefix(username string) string {
	return fmt.Sprintf("%s%s:", UserTokenPrefix, username)
}

// getTokenID 从token生成唯一标识
// 使用token的SHA256哈希值的前16位，确保唯一性
func (jwtService *JwtService) getTokenID(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])[:16]
}

// AddToken 添加一个新的token
// 单设备模式：覆盖旧token
// 多设备模式：添加新token，如果设置了限制，会删除最旧的token
func (jwtService *JwtService) AddToken(username string, token string, expiration time.Duration) error {
	ctx := context.Background()
	tokenKey := jwtService.generateTokenKey(username, jwtService.getTokenID(token))

	global.TD27_LOG.Debug("添加token",
		"username", username,
		"tokenKey", tokenKey,
		"expiration", expiration,
		"multiLogin", jwtService.isMultiLogin(),
		"limit", jwtService.getMultiLoginLimit())

	if jwtService.isMultiLogin() {
		// 多设备登录模式
		limit := jwtService.getMultiLoginLimit()
		if limit > 0 {
			// 获取该用户的所有token
			userTokens, err := jwtService.getUserTokens(ctx, username)
			if err != nil {
				global.TD27_LOG.Warn("获取用户token列表失败", "username", username, "error", err)
			}

			// 如果已达到限制，删除最旧的token
			if len(userTokens) >= limit {
				// 按过期时间排序，删除最旧的（过期时间最早的）
				sort.Slice(userTokens, func(i, j int) bool {
					return userTokens[i].ExpiresAt.Before(userTokens[j].ExpiresAt)
				})

				// 删除最旧的token
				tokensToRemove := len(userTokens) - limit + 1
				for i := 0; i < tokensToRemove && i < len(userTokens); i++ {
					if err = jwtService.cache.Del(ctx, userTokens[i].Key); err != nil {
						global.TD27_LOG.Warn("删除旧token失败",
							"key", userTokens[i].Key,
							"error", err)
					} else {
						global.TD27_LOG.Info("删除旧token",
							"username", username,
							"key", userTokens[i].Key)
					}
				}
			}
		}
	} else {
		// 单设备登录模式：直接覆盖旧token
		// 单设备模式下key不包含tokenID，直接Set即可覆盖
	}

	// 存储token，设置过期时间
	err := jwtService.cache.Set(ctx, username, tokenKey, token, expiration)
	if err != nil {
		global.TD27_LOG.Error("存储token失败",
			"username", username,
			"tokenKey", tokenKey,
			"error", err)
	} else {
		global.TD27_LOG.Debug("存储token成功",
			"username", username,
			"tokenKey", tokenKey)
	}
	return err
}

// userTokenInfo 用户token信息
type userTokenInfo struct {
	Key       string
	Token     string
	ExpiresAt time.Time
}

// getUserTokens 获取用户的所有token（多设备模式）
func (jwtService *JwtService) getUserTokens(ctx context.Context, username string) ([]userTokenInfo, error) {
	prefix := jwtService.getUserTokensPrefix(username)
	keys, err := jwtService.cache.ListKeysByPrefix(ctx, prefix)
	if err != nil {
		return nil, err
	}

	var tokens []userTokenInfo
	for _, key := range keys {
		token, err := jwtService.cache.Get(ctx, key)
		if err != nil {
			continue // 跳过过期或无效的token
		}

		ttl, err := jwtService.cache.TTL(ctx, key)
		if err != nil {
			continue
		}

		tokens = append(tokens, userTokenInfo{
			Key:       key,
			Token:     token,
			ExpiresAt: time.Now().Add(ttl),
		})
	}

	return tokens, nil
}

// ValidateToken 验证token是否有效
// 单设备模式：检查token是否匹配存储的值
// 多设备模式：检查token是否存在于用户的token列表中
func (jwtService *JwtService) ValidateToken(username string, token string) bool {
	ctx := context.Background()

	tokenKey := jwtService.generateTokenKey(username, jwtService.getTokenID(token))
	global.TD27_LOG.Debug("验证token",
		"username", username,
		"tokenKey", tokenKey,
		"multiLogin", jwtService.isMultiLogin())

	var err error
	var cachedToken string

	if jwtService.isMultiLogin() {
		// 多设备模式：检查特定token
		cachedToken, err = jwtService.cache.Get(ctx, tokenKey)
		if err != nil {
			global.TD27_LOG.Error("Get multi-login cacheToken failed",
				"username", username,
				"error", err)
			return false
		}
	} else {
		// 单设备模式：检查token是否匹配
		tokenKey = jwtService.generateTokenKey(username, "")
		cachedToken, err = jwtService.cache.Get(ctx, tokenKey)
		if err != nil {
			global.TD27_LOG.Error("Get single-login cacheToken failed",
				"username", username,
				"error", err)
			return false
		}
	}

	valid := cachedToken == token

	global.TD27_LOG.Debug("Token验证结果",
		"username", username,
		"tokenKey", tokenKey,
		"valid", valid)

	return valid
}

// RemoveToken 移除用户的某个token（用于登出）
func (jwtService *JwtService) RemoveToken(username string, token string) error {
	ctx := context.Background()

	if jwtService.isMultiLogin() {
		tokenKey := jwtService.generateTokenKey(username, jwtService.getTokenID(token))
		return jwtService.cache.Del(ctx, tokenKey)
	}

	// 单设备模式：直接删除用户的token
	tokenKey := jwtService.generateTokenKey(username, "")
	return jwtService.cache.Del(ctx, tokenKey)
}

// RemoveAllTokens 移除用户的所有token（用于强制下线）
func (jwtService *JwtService) RemoveAllTokens(username string) error {
	ctx := context.Background()

	// 单设备模式：直接删除
	if !jwtService.isMultiLogin() {
		tokenKey := jwtService.generateTokenKey(username, "")
		return jwtService.cache.Del(ctx, tokenKey)
	}

	// 多设备模式：查找并删除所有token
	prefix := jwtService.getUserTokensPrefix(username)
	keys, err := jwtService.cache.ListKeysByPrefix(ctx, prefix)
	if err != nil {
		return fmt.Errorf("获取用户token列表失败: %w", err)
	}

	if len(keys) == 0 {
		return nil
	}

	// 删除所有token
	if err = jwtService.cache.Del(ctx, keys...); err != nil {
		return fmt.Errorf("删除用户token失败: %w", err)
	}

	global.TD27_LOG.Info("移除用户所有token",
		"username", username,
		"count", len(keys))

	return nil
}

// GetUserActiveSessions 获取用户活跃的会话数量（多设备模式）
func (jwtService *JwtService) GetUserActiveSessions(username string) int {
	if !jwtService.isMultiLogin() {
		// 单设备模式：检查是否有token
		ctx := context.Background()
		tokenKey := jwtService.generateTokenKey(username, "")
		_, err := jwtService.cache.Get(ctx, tokenKey)
		if err != nil {
			return 0
		}
		return 1
	}

	// 多设备模式：统计token数量
	ctx := context.Background()
	prefix := jwtService.getUserTokensPrefix(username)
	keys, err := jwtService.cache.ListKeysByPrefix(ctx, prefix)
	if err != nil {
		return 0
	}
	return len(keys)
}

// GetCachedUser 从缓存获取用户信息
func (jwtService *JwtService) GetCachedUser(userID uint) (*sysManagement.UserModel, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s%d", UserCachePrefix, userID)
	data, err := jwtService.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var user sysManagement.UserModel
	if err = json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUserCache 删除用户缓存
func (jwtService *JwtService) DeleteUserCache(ctx context.Context, username string) error {
	return jwtService.cache.DelByUsername(ctx, username)
}
