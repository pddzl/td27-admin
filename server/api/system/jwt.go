package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/system"
)

type JwtApi struct{}

// JoinInBlacklist
// @Tags      Jwt
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "jwt加入黑名单"
// @Router    /jwt/jsonInBlacklist [post]
func (j *JwtApi) JoinInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JoinInBlacklist(jwt)
	if err != nil {
		global.TD27_LOG.Error("jwt作废失败", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	response.OkWithMessage("jwt作废成功", c)
}
