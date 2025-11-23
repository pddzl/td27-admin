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

func (mr *MenuRouter) InitMenuRouter(rg *gin.RouterGroup) {
	base := rg.Group("menu")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("addMenu", mr.menuApi.AddMenu)
	record.POST("editMenu", mr.menuApi.EditMenu)
	record.POST("deleteMenu", mr.menuApi.DeleteMenu)
	record.POST("getElTreeMenus", mr.menuApi.GetElTreeMenus)
	// without record
	base.GET("getMenus", mr.menuApi.GetMenus)
}
