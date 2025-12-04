package authority

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/authority"
	"server/internal/middleware"
)

type RoleRouter struct {
	roleApi *authority.RoleApi
}

func NewRoleRouter() *RoleRouter {
	return &RoleRouter{roleApi: authority.NewRoleApi()}
}

func (rr *RoleRouter) InitRoleRouter(rg *gin.RouterGroup) {
	base := rg.Group("role")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", rr.roleApi.Create)
	record.POST("delete", rr.roleApi.Delete)
	record.POST("update", rr.roleApi.Update)
	record.POST("editRoleMenu", rr.roleApi.EditRoleMenu)
	// without record
	base.POST("list", rr.roleApi.List)
}
