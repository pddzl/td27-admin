package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemReq "server/model/system/request"
	systemRep "server/model/system/response"
	"server/utils"
)

type MenuApi struct{}

func (ma *MenuApi) GetMenus(c *gin.Context) {
	userInfo, err := utils.GetUserInfo(c)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	list, err := menuService.GetMenus(userInfo.ID)
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
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func (ma *MenuApi) EditMenu(c *gin.Context) {
	var editMenuReq systemReq.EditMenuReq
	_ = c.ShouldBindJSON(&editMenuReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&editMenuReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := menuService.EditMenu(editMenuReq); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

func (ma *MenuApi) DeleteMenu(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := menuService.DeleteMenu(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (ma *MenuApi) GetElTreeMenus(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	list, ids, err := menuService.GetElTreeMenus(cId.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		response.OkWithDetailed(systemRep.Menu{
			List:    list,
			MenuIds: ids,
		}, "获取成功", c)
	}
}
