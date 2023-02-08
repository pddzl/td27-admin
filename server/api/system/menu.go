package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	systemReq "server/model/system/request"
	"server/utils"
)

type MenuApi struct{}

func (ma *MenuApi) GetMenus(c *gin.Context) {
	userInfo, err := utils.GetUserInfo(c)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	list, err := menuService.GetMenus(userInfo.Roles)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

func (ma *MenuApi) AddMenu(c *gin.Context) {
	var menuReq systemReq.Menu
	_ = c.ShouldBindJSON(&menuReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&menuReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if ok := menuService.AddMenu(menuReq); !ok {
		response.OkWithMessage("添加成功", c)
	} else {
		response.FailWithMessage("添加失败", c)
	}
}
