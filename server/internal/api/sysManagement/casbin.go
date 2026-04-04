package sysManagement

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type CasbinApi struct {
	casbinService *serviceSysManagement.CasbinService
	apiService    *serviceSysManagement.ApiService
}

func NewCasbinApi() *CasbinApi {
	return &CasbinApi{
		casbinService: serviceSysManagement.NewCasbinService(),
		apiService:    serviceSysManagement.NewApiService(),
	}
}

// Update
// @Tags      CasbinApi
// @Summary   更新角色API权限（使用统一权限表）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateRoleAPIReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /casbin/update [post]
func (ca *CasbinApi) Update(c *gin.Context) {
	var req modelSysManagement.UpdateRoleAPIReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ca.apiService.UpdateRoleAPIPermissions(req.RoleID, req.APIPermissionIDs); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("更新角色API权限失败", zap.Error(err))
	} else {
		common.OkWithMessage("更新成功", c)
	}
}


