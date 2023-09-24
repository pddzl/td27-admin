package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/global"
	"server/model/common/response"
	systemReq "server/model/system/request"
)

type CasbinApi struct{}

// EditCasbin
// @Tags      CasbinApi
// @Summary   编辑casbin
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.ReqCasbin true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /casbin/editCasbin [post]
func (ca *CasbinApi) EditCasbin(c *gin.Context) {
	var reqCasbin systemReq.ReqCasbin
	_ = c.ShouldBindJSON(&reqCasbin)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&reqCasbin); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := casbinService.EditCasbin(reqCasbin.RoleId, reqCasbin.CasbinInfos); err != nil {
		response.FailWithMessage("更新失败", c)
		global.TD27_LOG.Error("更新失败", zap.Error(err))
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
