package sysManagement

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type CasbinApi struct {
	casbinService *serviceSysManagement.CasbinService
}

func NewCasbinApi() *CasbinApi {
	return &CasbinApi{casbinService: serviceSysManagement.NewCasbinService()}
}

// EditCasbin
// @Tags      CasbinApi
// @Summary   编辑casbin
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.ReqCasbin true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /casbin/editCasbin [post]
func (ca *CasbinApi) EditCasbin(c *gin.Context) {
	var reqCasbin modelSysManagement.ReqCasbin
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
