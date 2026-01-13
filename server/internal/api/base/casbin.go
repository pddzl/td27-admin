package base

import (
	authorityReq "server/internal/model/base/request"
	"server/internal/model/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/service/base"
)

type CasbinApi struct {
	casbinService *base.CasbinService
}

func NewCasbinApi() *CasbinApi {
	return &CasbinApi{casbinService: base.NewCasbinService()}
}

// EditCasbin
// @Tags      CasbinApi
// @Summary   编辑casbin
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.ReqCasbin true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /casbin/editCasbin [post]
func (ca *CasbinApi) EditCasbin(c *gin.Context) {
	var reqCasbin authorityReq.ReqCasbin
	if err := c.ShouldBindJSON(&reqCasbin); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ca.casbinService.EditCasbin(reqCasbin.RoleId, reqCasbin.CasbinInfos); err != nil {
		common.Fail(c)
		global.TD27_LOG.Error("更新失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}
