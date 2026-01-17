package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type RoleRouter struct {
	roleApi *sysManagement.RoleApi
}

func NewRoleRouter() *RoleRouter {
	return &RoleRouter{roleApi: sysManagement.NewRoleApi()}
}

func (r *RoleRouter) InitRoleRouter(rg *gin.RouterGroup) {
	base := rg.Group("role")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", r.roleApi.Create)
	record.POST("delete", r.roleApi.Delete)
	record.POST("update", r.roleApi.Update)
	record.POST("updateRoleMenu", r.roleApi.UpdateRoleMenu)
	// without record
	base.POST("list", r.roleApi.List)
}
