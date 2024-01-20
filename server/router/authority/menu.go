package authority

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type MenuRouter struct{}

func (u *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuWithoutRouter := Router.Group("menu")

	menuApi := api.ApiGroupApp.Authority.MenuApi
	{
		menuRouter.POST("addMenu", menuApi.AddMenu)
		menuRouter.POST("editMenu", menuApi.EditMenu)
		menuRouter.POST("deleteMenu", menuApi.DeleteMenu)
		menuRouter.POST("getElTreeMenus", menuApi.GetElTreeMenus)
	}
	{
		menuWithoutRouter.GET("getMenus", menuApi.GetMenus)
	}
}
