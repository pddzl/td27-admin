package authority

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*UserRouter
	*RoleRouter
	*MenuRouter
	*ApiRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		UserRouter: NewUserRouter(),
		RoleRouter: NewRoleRouter(),
		MenuRouter: NewMenuRouter(),
		ApiRouter:  NewApiRouter(),
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitApiRouter(group)
	rg.InitMenuRouter(group)
	rg.InitRoleRouter(group)
	rg.InitUserRouter(group)
}
