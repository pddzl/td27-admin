package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	commonReq "server/internal/model/common/request"
	commonResp "server/internal/model/common/response"
	_ "server/internal/model/entity/authority"
	modelAuthority "server/internal/model/entity/authority"
	authorityReq "server/internal/model/entity/authority/request"
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
		commonResp.FailWithMessage(err.Error(), c)
	} else {
		commonResp.OkWithDetailed(list, "获取成功", c)
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if role, err := ra.roleService.Create(&roleModel); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("添加角色失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(role, "添加角色成功", c)
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
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.Delete(cId.ID); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("删除角色失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
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
	var eRole authorityReq.EditRole
	if err := c.ShouldBindJSON(&eRole); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ra.roleService.Update(eRole); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
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
	var editRE authorityReq.EditRoleMenu
	if err := c.ShouldBindJSON(&editRE); err != nil {
		commonResp.FailWithMessage(err.Error(), c)
		return
	}

	if err := ra.roleService.EditRoleMenu(editRE.RoleId, editRE.Ids); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
	}
}
