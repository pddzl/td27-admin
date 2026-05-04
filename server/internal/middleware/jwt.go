package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"server/internal/global"
	"server/internal/model/common"
	pkgJwt "server/internal/pkg/jwt"
	"server/internal/service/sysManagement"
	serviceSysTool "server/internal/service/sysTool"
)

var (
	jwtService          = sysManagement.NewJwtService()
	serviceTokenService = serviceSysTool.NewServiceTokenService()
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 首先尝试JWT认证 (x-token header)
		token := c.Request.Header.Get("x-token")
		if token != "" {
			j := pkgJwt.NewJWT()
			claims, err := j.ParseToken(token)
			if err == nil {
				// 验证token是否有效
				global.TD27_LOG.Debug("JWTAuth", "path", c.Request.URL.Path, "username", claims.Username)
				if jwtService.ValidateToken(claims.Username, token) {
					// Token即将过期，刷新
					if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
						claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.TD27_CONFIG.JWT.ExpiresTime) * time.Second))
						newToken, _ := j.CreateTokenByOldToken(token, *claims)
						newClaims, _ := j.ParseToken(newToken)
						c.Header("new-token", newToken)
						c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
						jwtService.RemoveToken(claims.Username, token)
						jwtService.AddToken(claims.Username, newToken, time.Duration(global.TD27_CONFIG.JWT.ExpiresTime)*time.Second)
					}
					c.Set("claims", claims)
					c.Next()
					return
				}
			}
			// JWT解析失败或无效，继续检查服务令牌
		}

		// 2. 尝试服务令牌认证 (X-Service-Token header) - 仅认证，不检查权限
		serviceToken := c.Request.Header.Get("X-Service-Token")
		if serviceToken != "" && strings.HasPrefix(serviceToken, "sk-") {
			tokenID, err := serviceTokenService.AuthenticateToken(serviceToken)
			if err != nil {
				common.FailWithDetailed(gin.H{}, "无效的访问令牌", c)
				c.Abort()
				return
			}

			// 认证通过，设置上下文信息，权限检查交给CasbinHandler
			c.Set("isServiceToken", true)
			c.Set("serviceTokenID", tokenID)
			c.Set("serviceTokenPrefix", serviceToken[:12])
			c.Next()
			return
		}

		// 3. 都没有认证通过
		common.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
		c.Abort()
	}
}
