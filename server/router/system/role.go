package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleRouter := Router.Group("role")
	roleApi := api.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("getRoles", roleApi.GetRoles)
		roleRouter.POST("addRole", roleApi.AddRole)
		roleRouter.DELETE("deleteRole", roleApi.DeleteRole)
		roleRouter.POST("editRole", roleApi.EditRole)
		roleRouter.POST("getRoleMenus", roleApi.GetRoleMenus)
	}
	return roleRouter
}
