package sysTool

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*CronRouter
	*FileRouter
	*ServiceTokenRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		CronRouter:         NewCronRouter(),
		FileRouter:         NewFileRouter(),
		ServiceTokenRouter: &ServiceTokenRouter{},
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitCronRouter(group)
	rg.InitFileRouter(group)
	rg.InitServiceTokenRouter(group)
}
