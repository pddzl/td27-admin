package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type RolePermissionApi struct {
	rolePermissionService *serviceSysManagement.RolePermissionService
}

func NewRolePermissionApi() *RolePermissionApi {
	return &RolePermissionApi{
		rolePermissionService: serviceSysManagement.NewRolePermissionService(),
	}
}

// Rebuild
// @Tags      RolePermissionApi
// @Summary   编辑用户菜单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.RebuildRolePermissionReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /role_permission/rebuild [post]
func (a *RolePermissionApi) Rebuild(c *gin.Context) {
	var req modelSysManagement.RebuildRolePermissionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}

	if err := a.rolePermissionService.Rebuild(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", "error", err)
	} else {
		common.Ok(c)
	}
}
