package authority

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*LogRegRouter
	*CasbinRouter
	*UserRouter
	*RoleRouter
	*MenuRouter
	*ApiRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		LogRegRouter: NewLogRegRouter(),
		CasbinRouter: NewCasbinRouter(),
		UserRouter:   NewUserRouter(),
		RoleRouter:   NewRoleRouter(),
		MenuRouter:   NewMenuRouter(),
		ApiRouter:    NewApiRouter(),
	}
}

func (r *RouterGroup) InitPublic(group *gin.RouterGroup) {
	r.InitLogRegRouter(group)
}

func (r *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	r.InitCasbinRouter(group)
	r.InitUserRouter(group)
	r.InitRoleRouter(group)
	r.InitMenuRouter(group)
	r.InitApiRouter(group)
}
