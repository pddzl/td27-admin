package authority

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/global"
	authorityReq "server/model/authority/request"
	authorityRes "server/model/authority/response"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
	"server/utils"
)

type MenuApi struct{}

// GetMenus
// @Tags      MenuApi
// @Summary   获取用户菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]authority.MenuModel,msg=string}
// @Router    /menu/getMenus [post]
func (ma *MenuApi) GetMenus(c *gin.Context) {
	userInfo, err := utils.GetUserInfo(c)
	if err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	list, err := menuService.GetMenus(userInfo.ID)
	if err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(list, "获取成功", c)
	}
}

// AddMenu
// @Tags      MenuApi
// @Summary   添加菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditMenuReq true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /menu/addMenu [post]
func (ma *MenuApi) AddMenu(c *gin.Context) {
	var menuReq authorityReq.Menu
	_ = c.ShouldBindJSON(&menuReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&menuReq); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if ok := menuService.AddMenu(menuReq); !ok {
		commonRes.FailWithMessage("添加失败", c)
	} else {
		commonRes.OkWithMessage("添加成功", c)
	}
}

// EditMenu
// @Tags      MenuApi
// @Summary   编辑菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditMenuReq true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /menu/editMenu [post]
func (ma *MenuApi) EditMenu(c *gin.Context) {
	var editMenuReq authorityReq.EditMenuReq
	_ = c.ShouldBindJSON(&editMenuReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&editMenuReq); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := menuService.EditMenu(editMenuReq); err != nil {
		commonRes.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("编辑成功", c)
	}
}

// DeleteMenu
// @Tags      MenuApi
// @Summary   删除菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /menu/deleteMenu [post]
func (ma *MenuApi) DeleteMenu(c *gin.Context) {
	var cId commonReq.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := menuService.DeleteMenu(cId.ID); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}

// GetElTreeMenus
// @Tags      MenuApi
// @Summary   获取菜单树
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{data=authorityRes.Menu{list=[]authority.MenuModel,menuIds=[]uint},msg=string}
// @Router    /menu/getElTreeMenus [post]
func (ma *MenuApi) GetElTreeMenus(c *gin.Context) {
	var cId commonReq.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		commonRes.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	list, ids, err := menuService.GetElTreeMenus(cId.ID)
	if err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(authorityRes.Menu{
			List:    list,
			MenuIds: ids,
		}, "获取成功", c)
	}
}
