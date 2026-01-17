package sysManagement

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type RoleApi struct {
	roleService *serviceSysManagement.RoleService
}

func NewRoleApi() *RoleApi {
	return &RoleApi{roleService: serviceSysManagement.NewRoleService()}
}

// List
// @Tags      RoleApi
// @Summary   获取所有角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  common.Response{data=[]modelSysManagement.RoleModel,msg=string}
// @Router    /role/list [post]
func (ra *RoleApi) List(c *gin.Context) {
	var req common.PageInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := ra.roleService.List(&req); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("get roles failed", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

// Create
// @Tags      RoleApi
// @Summary   添加角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.RoleModel true "请求参数"
// @Success   200   {object}  common.Response{data=modelSysManagement.RoleModel,msg=string}
// @Router    /role/create [post]
func (ra *RoleApi) Create(c *gin.Context) {
	var req modelSysManagement.RoleModel
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := ra.roleService.Create(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "添加角色成功", c)
	}
}

// Delete
// @Tags      RoleApi
// @Summary   删除角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /role/delete [post]
func (ra *RoleApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.Delete(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("删除角色失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

// Update
// @Tags      RoleApi
// @Summary   编辑角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateRoleReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /role/update [post]
func (ra *RoleApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateRoleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

// UpdateRoleMenu
// @Tags      UpdateRoleMenu
// @Summary   编辑用户菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateRoleMenuReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /role/updateRoleMenu [post]
func (ra *RoleApi) UpdateRoleMenu(c *gin.Context) {
	var req modelSysManagement.UpdateRoleMenuReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	if err := ra.roleService.UpdateRoleMenu(req.RoleId, req.Ids); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}
