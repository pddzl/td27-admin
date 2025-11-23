package sysSet

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*DictRouter
	*DictDetailRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		DictRouter:       NewDictRouter(),
		DictDetailRouter: NewDictDetailRouter(),
	}
}

func (rg *RouterGroup) InitPublic(group *gin.RouterGroup) {}

func (rg *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	rg.InitDictRouter(group)
	rg.InitDictDetailRouter(group)
}
