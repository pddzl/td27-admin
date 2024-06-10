package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	modelAuthority "server/model/authority"
	authorityReq "server/model/authority/request"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
)

type RoleApi struct{}

// GetRoles
// @Tags      RoleApi
// @Summary   获取所有角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]authority.RoleModel,msg=string}
// @Router    /role/getRoles [post]
func (ra *RoleApi) GetRoles(c *gin.Context) {
	if list, err := roleService.GetRoles(); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
	} else {
		commonRes.OkWithDetailed(list, "获取成功", c)
	}
}

// AddRole
// @Tags      RoleApi
// @Summary   添加角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelAuthority.RoleModel true "请求参数"
// @Success   200   {object}  response.Response{data=authority.RoleModel,msg=string}
// @Router    /api/addRole [post]
func (ra *RoleApi) AddRole(c *gin.Context) {
	var roleModel modelAuthority.RoleModel
	if err := c.ShouldBindJSON(&roleModel); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if role, err := roleService.AddRole(&roleModel); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(role, "添加角色成功", c)
	}
}

// DeleteRole
// @Tags      RoleApi
// @Summary   删除角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/deleteRole [post]
func (ra *RoleApi) DeleteRole(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := roleService.DeleteRole(cId.ID); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("删除角色失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

// EditRole
// @Tags      RoleApi
// @Summary   编辑角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditRole true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editRole [post]
func (ra *RoleApi) EditRole(c *gin.Context) {
	var eRole authorityReq.EditRole
	if err := c.ShouldBindJSON(&eRole); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	if err := roleService.EditRole(eRole); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}

// EditRoleMenu
// @Tags      RoleApi
// @Summary   编辑用户菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditRoleMenu true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editRoleMenu [post]
func (ra *RoleApi) EditRoleMenu(c *gin.Context) {
	var editRE authorityReq.EditRoleMenu
	if err := c.ShouldBindJSON(&editRE); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if err := roleService.EditRoleMenu(editRE.RoleId, editRE.Ids); err != nil {
		commonRes.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.Ok(c)
	}
}
