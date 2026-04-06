package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type RolePermissionRouter struct {
	rolePermissionApi *sysManagement.RolePermissionApi
}

func NewRolePermissionRouter() *RolePermissionRouter {
	return &RolePermissionRouter{rolePermissionApi: sysManagement.NewRolePermissionApi()}
}

func (r *RolePermissionRouter) InitRolePermissionRouter(rg *gin.RouterGroup) {
	base := rg.Group("role_permission")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("update", r.rolePermissionApi.Update)
}
