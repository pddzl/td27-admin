package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type MenuRouter struct{}

func (u *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu")
	menuApi := api.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.GET("getMenus", menuApi.GetMenus)
		menuRouter.POST("addMenu", menuApi.AddMenu)
		menuRouter.POST("editMenu", menuApi.EditMenu)
		menuRouter.POST("deleteMenu", menuApi.DeleteMenu)
		menuRouter.POST("getElTreeMenus", menuApi.GetElTreeMenus)
	}
	return menuRouter
}
