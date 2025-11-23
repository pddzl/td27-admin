package sysTool

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*CronRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{CronRouter: NewCronRouter()}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitCronRouter(group)
}
