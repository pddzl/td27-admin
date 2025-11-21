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

func (rr *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	// record
	roleRouter := Router.Group("role").Use(middleware.OperationRecord())
	roleRouter.POST("addRole", rr.roleApi.AddRole)
	roleRouter.POST("deleteRole", rr.roleApi.DeleteRole)
	roleRouter.POST("editRole", rr.roleApi.EditRole)
	roleRouter.POST("editRoleMenu", rr.roleApi.EditRoleMenu)
	// without record
	roleWithoutRouter := Router.Group("role")
	roleWithoutRouter.POST("getRoles", rr.roleApi.GetRoles)
}
