package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/authority/role"
	"server/internal/model/common"
	serviceAuthority "server/internal/service/authority"
)

type RoleApi struct {
	roleService *serviceAuthority.RoleService
}

func NewRoleApi() *RoleApi {
	return &RoleApi{roleService: serviceAuthority.NewRoleService()}
}

// List
// @Tags      RoleApi
// @Summary   获取所有角色
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  common.Response{data=[]role.RoleModel,msg=string}
// @Router    /role/getRoles [post]
func (ra *RoleApi) List(c *gin.Context) {
	var pageInfo common.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if list, total, err := ra.roleService.List(&pageInfo); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("get roles failed", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
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
// @Param     data  body      role.RoleModel true "请求参数"
// @Success   200   {object}  common.Response{data=role.RoleModel,msg=string}
// @Router    /api/addRole [post]
func (ra *RoleApi) Create(c *gin.Context) {
	var roleModel role.RoleModel
	if err := c.ShouldBindJSON(&roleModel); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := ra.roleService.Create(&roleModel); err != nil {
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
// @Router    /api/deleteRole [post]
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
// @Param     data  body      role.UpdateRoleReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /api/updateRole [post]
func (ra *RoleApi) Update(c *gin.Context) {
	var req role.UpdateRoleReq
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
// @Param     data  body      role.UpdateRoleMenuReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /api/editRoleMenu [post]
func (ra *RoleApi) UpdateRoleMenu(c *gin.Context) {
	var req role.UpdateRoleMenuReq
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
