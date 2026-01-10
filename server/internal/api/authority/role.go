package authority

import (
	"server/internal/model/authority/menu"
	modelAuthority "server/internal/model/authority/role"
	commonReq "server/internal/model/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	_ "server/internal/model/entity/authority"
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
// @Success   200   {object}  commonResp.Response{data=[]authority.RoleModel,msg=string}
// @Router    /role/getRoles [post]
func (ra *RoleApi) List(c *gin.Context) {
	if list, err := ra.roleService.List(); err != nil {
		commonReq.FailWithMessage(err.Error(), c)
	} else {
		commonReq.OkWithDetailed(list, "获取成功", c)
	}
}

// Create
// @Tags      RoleApi
// @Summary   添加角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelAuthority.RoleModel true "请求参数"
// @Success   200   {object}  commonResp.Response{data=authority.RoleModel,msg=string}
// @Router    /api/addRole [post]
func (ra *RoleApi) Create(c *gin.Context) {
	var roleModel modelAuthority.RoleModel
	if err := c.ShouldBindJSON(&roleModel); err != nil {
		commonReq.FailReq(err.Error(), c)
		return
	}

	if role, err := ra.roleService.Create(&roleModel); err != nil {
		commonReq.Fail(c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		commonReq.OkWithDetailed(role, "添加角色成功", c)
	}
}

// Delete
// @Tags      RoleApi
// @Summary   删除角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /api/deleteRole [post]
func (ra *RoleApi) Delete(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonReq.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.Delete(cId.ID); err != nil {
		commonReq.Fail(c)
		global.TD27_LOG.Error("删除角色失败", zap.Error(err))
	} else {
		commonReq.Ok(c)
	}
}

// Update
// @Tags      RoleApi
// @Summary   编辑角色
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditRole true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /api/update [post]
func (ra *RoleApi) Update(c *gin.Context) {
	var eRole modelAuthority.EditRole
	if err := c.ShouldBindJSON(&eRole); err != nil {
		commonReq.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.Update(eRole); err != nil {
		commonReq.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonReq.Ok(c)
	}
}

// EditRoleMenu
// @Tags      RoleApi
// @Summary   编辑用户菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditRoleMenu true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /api/editRoleMenu [post]
func (ra *RoleApi) EditRoleMenu(c *gin.Context) {
	var editRE menu.EditRoleMenu
	if err := c.ShouldBindJSON(&editRE); err != nil {
		commonReq.FailWithMessage(err.Error(), c)
		return
	}

	if err := ra.roleService.EditRoleMenu(editRE.RoleId, editRE.Ids); err != nil {
		commonReq.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonReq.Ok(c)
	}
}
