package fileM

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*FileRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		FileRouter: NewFileRouter(),
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitFileRouter(group)
}
