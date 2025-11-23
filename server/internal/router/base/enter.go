package base

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*CasbinRouter
	*LogRegRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		CasbinRouter: NewCasbinRouter(),
		LogRegRouter: NewLogRegRouter(),
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {
	rg.InitLogRegRouter(group)
}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitCasbinRouter(group)
}
