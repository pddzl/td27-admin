package base

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/global"
	authorityReq "server/model/authority/request"
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
	_ = c.ShouldBindJSON(&reqCasbin)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&reqCasbin); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := casbinService.EditCasbin(reqCasbin.RoleId, reqCasbin.CasbinInfos); err != nil {
		commonRes.FailWithMessage("更新失败", c)
		global.TD27_LOG.Error("更新失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("更新成功", c)
	}
}
