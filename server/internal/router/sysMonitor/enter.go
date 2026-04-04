package sysMonitor

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*OperationLogRouter
	*DashboardRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		OperationLogRouter: NewOperationLogRouter(),
		DashboardRouter:    NewDashboardRouter(),
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitOperationLogRouter(group)
	rg.InitDashboardRouter(group)
}
