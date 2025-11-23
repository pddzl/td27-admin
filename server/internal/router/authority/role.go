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
	record.POST("addRole", rr.roleApi.AddRole)
	record.POST("deleteRole", rr.roleApi.DeleteRole)
	record.POST("editRole", rr.roleApi.EditRole)
	record.POST("editRoleMenu", rr.roleApi.EditRoleMenu)
	// without record
	base.POST("getRoles", rr.roleApi.GetRoles)
}
