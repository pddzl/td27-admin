package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"

	"server/global"
	modelAuthority "server/model/authority"
	commonRes "server/model/common/response"
	"server/service"
	"server/utils"
)

var jwtService = service.ServiceGroupApp.Base.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			commonRes.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.TokenExpired) {
				commonRes.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			commonRes.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// x-token与redis存的token做对比
		redisJwtToken, err := jwtService.GetRedisJWT(claims.Username)
		if redisJwtToken != token {
			commonRes.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}

		// 用户是否存在
		var userModel modelAuthority.UserModel
		err = global.TD27_DB.Where("id = ?", claims.ID).First(&userModel).Error
		if err != nil {
			commonRes.FailWithMessage("用户不存在", c)
			c.Abort()
			global.TD27_LOG.Error("用户不存在")
			return
		}

		// 已登录用户是否禁用
		if !userModel.Active {
			commonRes.FailWithMessage("用户被禁用", c)
			c.Abort()
			global.TD27_LOG.Error("用户被禁用")
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(global.TD27_CONFIG.JWT.ExpiresTime) * time.Second))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
