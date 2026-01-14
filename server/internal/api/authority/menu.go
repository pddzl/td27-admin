package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	modelAuthority "server/internal/model/authority"
	"server/internal/model/common"
	serviceAuthority "server/internal/service/authority"
)

type MenuApi struct {
	menuService *serviceAuthority.MenuService
}

func NewMenuApi() *MenuApi {
	return &MenuApi{menuService: serviceAuthority.NewMenuService()}
}

// List
// @Tags      MenuApi
// @Summary   获取用户菜单
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  common.Response{data=[]modelAuthority.MenuModel,msg=string}
// @Router    /menu/list [post]
func (ma *MenuApi) List(c *gin.Context) {
	userInfo, err := GetUserInfo(c)
	if err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	list, err := ma.menuService.List(userInfo.ID)
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
// @Param     data  body      modelAuthority.Menu true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /menu/create [post]
func (ma *MenuApi) Create(c *gin.Context) {
	var req modelAuthority.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ma.menuService.Create(&req); err != nil {
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
// @Param     data  body      modelAuthority.UpdateMenuReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /menu/update [post]
func (ma *MenuApi) Update(c *gin.Context) {
	var req modelAuthority.UpdateMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ma.menuService.Update(&req); err != nil {
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
func (ma *MenuApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ma.menuService.Delete(cId.ID); err != nil {
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
// @Success   200   {object}  common.Response{data=modelAuthority.MenuResp{list=[]modelAuthority.MenuModel,menuIds=[]uint},msg=string}
// @Router    /menu/getElTreeMenus [post]
func (ma *MenuApi) GetElTreeMenus(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, ids, err := ma.menuService.GetElTreeMenus(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		common.OkWithDetailed(modelAuthority.MenuResp{
			List:    list,
			MenuIds: ids,
		}, "获取成功", c)
	}
}
