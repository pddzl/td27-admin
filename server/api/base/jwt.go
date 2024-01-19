package base

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	modelBase "server/model/base"
	commonRes "server/model/common/response"
)

type JwtApi struct{}

// JoinInBlacklist
// @Tags      JwtApi
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}
// @Router    /jwt/jsonInBlacklist [POST]
func (j *JwtApi) JoinInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := modelBase.JwtBlackListModel{Jwt: token}
	err := jwtService.JoinInBlacklist(jwt)
	if err != nil {
		global.TD27_LOG.Error("jwt作废失败", zap.Error(err))
		commonRes.FailWithMessage("jwt作废失败", c)
		return
	}
	commonRes.OkWithMessage("jwt作废成功", c)
}
