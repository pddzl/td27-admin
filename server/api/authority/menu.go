package authority

import (
	"github.com/gin-gonic/gin"
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
	if err := c.ShouldBindJSON(&menuReq); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if ok := menuService.AddMenu(menuReq); !ok {
		commonRes.Fail(c)
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
	if err := c.ShouldBindJSON(&editMenuReq); err != nil {
		commonRes.FailReq(err.Error(), c)
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
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
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
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if list, ids, err := menuService.GetElTreeMenus(cId.ID); err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(authorityRes.Menu{
			List:    list,
			MenuIds: ids,
		}, "获取成功", c)
	}
}
