package sysManagement

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type MenuApi struct {
	menuService *serviceSysManagement.MenuService
}

func NewMenuApi() *MenuApi {
	return &MenuApi{menuService: serviceSysManagement.NewMenuService()}
}

// List
// @Tags      MenuApi
// @Summary   获取用户菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  common.Response{data=[]modelSysManagement.MenuData,msg=string}
// @Router    /menu/list [get]
func (a *MenuApi) List(c *gin.Context) {
	userInfo, err := GetUserInfo(c)
	if err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
		return
	}

	list, err := a.menuService.List(userInfo.ID)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		common.OkWithDetailed(list, "获取成功", c)
	}
}

// Create
// @Tags      MenuApi
// @Summary   添加菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.Menu true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /menu/create [post]
func (a *MenuApi) Create(c *gin.Context) {
	var req modelSysManagement.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if _, err := a.menuService.Create(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
	} else {
		common.OkWithMessage("添加成功", c)
	}
}

// Update
// @Tags      MenuApi
// @Summary   编辑菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateMenuReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /menu/update [post]
func (a *MenuApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.menuService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		common.OkWithMessage("编辑成功", c)
	}
}

// Delete
// @Tags      MenuApi
// @Summary   删除菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /menu/delete [post]
func (a *MenuApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := a.menuService.Delete(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.OkWithMessage("删除成功", c)
	}
}

// GetElTreeMenus
// @Tags      MenuApi
// @Summary   获取菜单树
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true "请求参数"
// @Success   200   {object}  common.Response{data=modelSysManagement.MenuResp{list=[]modelSysManagement.MenuModel,menuIds=[]uint},msg=string}
// @Router    /menu/getElTreeMenus [post]
func (a *MenuApi) GetElTreeMenus(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, ids, err := a.menuService.GetElTreeMenus(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		common.OkWithDetailed(modelSysManagement.MenuResp{
			List:    list,
			MenuIds: ids,
		}, "获取成功", c)
	}
}
