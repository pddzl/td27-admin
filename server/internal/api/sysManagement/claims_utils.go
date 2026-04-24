package sysManagement

import (
	"github.com/gin-gonic/gin"

	modelSysManagement "server/internal/model/sysManagement"
	pkgJwt "server/internal/pkg/jwt"
	"log/slog"
)

func GetUserInfo(c *gin.Context) (*modelSysManagement.CustomClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		token := c.Request.Header.Get("x-token")
		j := pkgJwt.NewJWT()
		claimsParse, err := j.ParseToken(token)
		if err != nil {
			slog.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
			return nil, err
		}
		return claimsParse, nil
	}

	return claims.(*modelSysManagement.CustomClaims), nil
}

func GetClaims(c *gin.Context) (*modelSysManagement.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := pkgJwt.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		slog.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
