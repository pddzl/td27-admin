package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	commonReq "server/internal/model/common/request"
	commonResp "server/internal/model/common/response"
	_ "server/internal/model/entity/authority"
	authorityReq "server/internal/model/entity/authority/request"
	authorityResp "server/internal/model/entity/authority/response"
	"server/internal/pkg"
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
// @Success   200   {object}  commonResp.Response{data=[]authority.MenuModel,msg=string}
// @Router    /menu/list [post]
func (ma *MenuApi) List(c *gin.Context) {
	userInfo, err := pkg.GetUserInfo(c)
	if err != nil {
		commonResp.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	list, err := ma.menuService.List(userInfo.ID)
	if err != nil {
		commonResp.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(list, "获取成功", c)
	}
}

// Create
// @Tags      MenuApi
// @Summary   添加菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditMenuReq true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /menu/create [post]
func (ma *MenuApi) Create(c *gin.Context) {
	var menuReq authorityReq.Menu
	if err := c.ShouldBindJSON(&menuReq); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if ok := ma.menuService.Create(menuReq); !ok {
		commonResp.Fail(c)
	} else {
		commonResp.OkWithMessage("添加成功", c)
	}
}

// Update
// @Tags      MenuApi
// @Summary   编辑菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditMenuReq true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /menu/update [post]
func (ma *MenuApi) Update(c *gin.Context) {
	var editMenuReq authorityReq.EditMenuReq
	if err := c.ShouldBindJSON(&editMenuReq); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ma.menuService.Update(editMenuReq); err != nil {
		commonResp.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonResp.OkWithMessage("编辑成功", c)
	}
}

// Delete
// @Tags      MenuApi
// @Summary   删除菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /menu/delete [post]
func (ma *MenuApi) Delete(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ma.menuService.Delete(cId.ID); err != nil {
		commonResp.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonResp.OkWithMessage("删除成功", c)
	}
}

// GetElTreeMenus
// @Tags      MenuApi
// @Summary   获取菜单树
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{data=authorityResp.Menu{list=[]authority.MenuModel,menuIds=[]uint},msg=string}
// @Router    /menu/getElTreeMenus [post]
func (ma *MenuApi) GetElTreeMenus(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if list, ids, err := ma.menuService.GetElTreeMenus(cId.ID); err != nil {
		commonResp.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(authorityResp.Menu{
			List:    list,
			MenuIds: ids,
		}, "获取成功", c)
	}
}
