package authority

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role").Use(middleware.OperationRecord())
	roleWithoutRouter := Router.Group("role")

	roleApi := api.ApiGroupApp.Authority.RoleApi
	{
		roleRouter.POST("addRole", roleApi.AddRole)
		roleRouter.POST("deleteRole", roleApi.DeleteRole)
		roleRouter.POST("editRole", roleApi.EditRole)
		roleRouter.POST("editRoleMenu", roleApi.EditRoleMenu)
	}
	{
		roleWithoutRouter.POST("getRoles", roleApi.GetRoles)
	}
}
