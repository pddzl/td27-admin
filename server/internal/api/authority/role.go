package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common/request"
	"server/internal/model/common/response"
	modelAuthority "server/internal/model/entity/authority"
	authorityReq "server/internal/model/entity/authority/request"
	"server/internal/service/authority"
)

type RoleApi struct {
	roleService *authority.RoleService
}

func NewRoleApi() *RoleApi {
	return &RoleApi{roleService: authority.NewRoleService()}
}

// GetRoles
// @Tags      RoleApi
// @Summary   获取所有角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]authority.RoleModel,msg=string}
// @Router    /role/getRoles [post]
func (ra *RoleApi) GetRoles(c *gin.Context) {
	if list, err := ra.roleService.GetRoles(); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
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
		response.FailReq(err.Error(), c)
		return
	}

	if role, err := ra.roleService.AddRole(&roleModel); err != nil {
		response.Fail(c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		response.OkWithDetailed(role, "添加角色成功", c)
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
	var cId request.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		response.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.DeleteRole(cId.ID); err != nil {
		response.Fail(c)
		global.TD27_LOG.Error("删除角色失败", zap.Error(err))
	} else {
		response.Ok(c)
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
		response.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.EditRole(eRole); err != nil {
		response.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.Ok(c)
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
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ra.roleService.EditRoleMenu(editRE.RoleId, editRE.Ids); err != nil {
		response.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.Ok(c)
	}
}
