package sysManagement

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*LogRegRouter
	*CasbinRouter
	*UserRouter
	*RoleRouter
	*RolePermissionRouter
	*MenuRouter
	*ApiRouter
	*DictRouter
	*DictDetailRouter
	*DeptRouter
}

func NewRouterGroup() *RouterGroup {
	return &RouterGroup{
		LogRegRouter:         NewLogRegRouter(),
		CasbinRouter:         NewCasbinRouter(),
		UserRouter:           NewUserRouter(),
		RoleRouter:           NewRoleRouter(),
		RolePermissionRouter: NewRolePermissionRouter(),
		MenuRouter:           NewMenuRouter(),
		ApiRouter:            NewApiRouter(),
		DictRouter:           NewDictRouter(),
		DictDetailRouter:     NewDictDetailRouter(),
		DeptRouter:           NewDeptRouter(),
	}
}

func (r *RouterGroup) InitPublic(group *gin.RouterGroup) {
	r.InitLogRegRouter(group)
}

func (r *RouterGroup) InitPrivate(group *gin.RouterGroup) {
	r.InitCasbinRouter(group)
	r.InitUserRouter(group)
	r.InitRoleRouter(group)
	r.InitRolePermissionRouter(group)
	r.InitMenuRouter(group)
	r.InitApiRouter(group)
	r.InitDictRouter(group)
	r.InitDictDetailRouter(group)
	r.InitDeptRouter(group)
}
