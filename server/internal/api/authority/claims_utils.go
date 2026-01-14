package authority

import (
	"github.com/gin-gonic/gin"

	"server/internal/global"
	modelAuthority "server/internal/model/authority"
	pkgJwt "server/internal/pkg/jwt"
)

func GetUserInfo(c *gin.Context) (*modelAuthority.CustomClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		token := c.Request.Header.Get("x-token")
		j := pkgJwt.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			global.TD27_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
			return nil, err
		}
		return claims, nil
	}

	return claims.(*modelAuthority.CustomClaims), nil
}

func GetClaims(c *gin.Context) (*modelAuthority.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := pkgJwt.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.TD27_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
