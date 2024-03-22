package base

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	authorityReq "server/model/base/request"
	commonRes "server/model/common/response"
)

type CasbinApi struct{}

// EditCasbin
// @Tags      CasbinApi
// @Summary   编辑casbin
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.ReqCasbin true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /casbin/editCasbin [post]
func (ca *CasbinApi) EditCasbin(c *gin.Context) {
	var reqCasbin authorityReq.ReqCasbin
	if err := c.ShouldBindJSON(&reqCasbin); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := casbinService.EditCasbin(reqCasbin.RoleId, reqCasbin.CasbinInfos); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("更新失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}
