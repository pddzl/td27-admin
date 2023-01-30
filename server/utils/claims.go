package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/system"
	systemReq "server/model/system/request"
)

func GetUserInfo(c *gin.Context) (*systemReq.CustomClaimsRole, error) {
	var claimsParse systemReq.CustomClaimsRole
	var userModel system.UserModel
	claims, exists := c.Get("claims")
	if !exists {
		token := c.Request.Header.Get("x-token")
		j := NewJWT()
		claimsRaw, err := j.ParseToken(token)
		claimsParse.CustomClaims = claimsRaw
		if err != nil {
			global.TD27_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
			return nil, err
		}
	} else {
		claimsRaw, ok := claims.(*systemReq.CustomClaims)
		if ok {
			claimsParse.CustomClaims = claimsRaw
		} else {
			global.TD27_LOG.Error("claims断言失败")
			return nil, errors.New("claims断言失败")
		}
	}

	// 获取用户roles
	err := global.TD27_DB.Model(&system.UserModel{}).Where("id = ?", claimsParse.CustomClaims.ID).Preload("Roles").Find(&userModel).Error
	if err != nil {
		global.TD27_LOG.Error("获取userInfo roles失败")
		return nil, err
	}

	for _, value := range userModel.Roles {
		claimsParse.Roles = append(claimsParse.Roles, value.RoleName)
	}

	return &claimsParse, err
}
