package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemReq "server/model/system/request"
)

type RoleApi struct{}

// GetRoles
// @Tags      RoleApi
// @Summary   获取所有角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]system.RoleModel,msg=string}
// @Router    /role/getRoles [post]
func (ra *RoleApi) GetRoles(c *gin.Context) {
	if list, err := roleService.GetRoles(); err != nil {
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
// @Param     data  body      systemReq.Role true "请求参数"
// @Success   200   {object}  response.Response{data=system.RoleModel,msg=string}
// @Router    /api/addRole [post]
func (ra *RoleApi) AddRole(c *gin.Context) {
	var roleReq systemReq.Role
	_ = c.ShouldBindJSON(&roleReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&roleReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if role, err := roleService.AddRole(roleReq.RoleName); err != nil {
		response.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		if err = casbinService.EditCasbin(role.ID, systemReq.DefaultCasbin()); err != nil {
			global.TD27_LOG.Error("更新casbin rule失败", zap.Error(err))
		}
		response.OkWithDetailed(role, "添加成功", c)
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
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := roleService.DeleteRole(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除角色失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// EditRole
// @Tags      RoleApi
// @Summary   编辑角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.EditRole true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editRole [post]
func (ra *RoleApi) EditRole(c *gin.Context) {
	var eRole systemReq.EditRole
	_ = c.ShouldBindJSON(&eRole)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&eRole); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := roleService.EditRole(eRole); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// EditRoleMenu
// @Tags      RoleApi
// @Summary   编辑用户菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.EditRoleMenu true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /api/editRoleMenu [post]
func (ra *RoleApi) EditRoleMenu(c *gin.Context) {
	var editRE systemReq.EditRoleMenu
	_ = c.ShouldBindJSON(&editRE)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&editRE); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := roleService.EditRoleMenu(editRE.RoleId, editRE.Ids); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}
