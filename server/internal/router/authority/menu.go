package authority

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/authority"
	"server/internal/middleware"
)

type MenuRouter struct {
	menuApi *authority.MenuApi
}

func NewMenuRouter() *MenuRouter {
	return &MenuRouter{menuApi: authority.NewMenuApi()}
}

func (mr *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	// record
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouter.POST("addMenu", mr.menuApi.AddMenu)
	menuRouter.POST("editMenu", mr.menuApi.EditMenu)
	menuRouter.POST("deleteMenu", mr.menuApi.DeleteMenu)
	menuRouter.POST("getElTreeMenus", mr.menuApi.GetElTreeMenus)
	// without record
	menuWithoutRouter := Router.Group("menu")
	menuWithoutRouter.GET("getMenus", mr.menuApi.GetMenus)
}
