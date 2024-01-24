package utils

import (
	"github.com/gin-gonic/gin"
	"server/global"
	baseReq "server/model/base/request"
)

func GetUserInfo(c *gin.Context) (*baseReq.CustomClaims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		token := c.Request.Header.Get("x-token")
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			global.TD27_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
			return nil, err
		}
		return claims, nil
	} else {
		return claims.(*baseReq.CustomClaims), nil
	}
}

func GetClaims(c *gin.Context) (*baseReq.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.TD27_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}
