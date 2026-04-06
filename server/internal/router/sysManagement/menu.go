package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type MenuRouter struct {
	menuApi *sysManagement.MenuApi
}

func NewMenuRouter() *MenuRouter {
	return &MenuRouter{menuApi: sysManagement.NewMenuApi()}
}

func (r *MenuRouter) InitMenuRouter(rg *gin.RouterGroup) {
	base := rg.Group("menu")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", r.menuApi.Create)
	record.POST("update", r.menuApi.Update)
	record.POST("delete", r.menuApi.Delete)
	record.POST("getElTreeMenus", r.menuApi.GetElTreeMenus)
	// without record
	base.GET("list", r.menuApi.List)
}
