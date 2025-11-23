package monitor

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*OperationLogRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		OperationLogRouter: NewOperationLogRouter(),
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitOperationLogRouter(group)
}
