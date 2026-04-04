package middleware

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"server/internal/global"
	"server/internal/model/common"
	pkgJwt "server/internal/pkg/jwt"
	"server/internal/service/sysManagement"
)

var (
	jwtService = sysManagement.NewJwtService()
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			common.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := pkgJwt.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, pkgJwt.TokenExpired) {
				common.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			common.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 验证token是否有效（支持单设备/多设备模式切换）
		if !jwtService.ValidateToken(claims.Username, token) {
			if global.TD27_CONFIG.JWT.MultiLogin {
				common.FailWithDetailed(gin.H{"reload": true}, "登录已过期，请重新登录", c)
			} else {
				common.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			}
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.TD27_CONFIG.JWT.ExpiresTime) * time.Second))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))

			// 更新缓存：删除旧token，存储新token
			jwtService.RemoveToken(claims.Username, token)
			jwtService.AddToken(claims.Username, newToken, time.Duration(global.TD27_CONFIG.JWT.ExpiresTime)*time.Second)
		}
		c.Set("claims", claims)
		c.Next()
	}
}
