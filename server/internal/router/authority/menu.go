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
	record.POST("create", mr.menuApi.Create)
	record.POST("update", mr.menuApi.Update)
	record.POST("delete", mr.menuApi.Delete)
	record.POST("getElTreeMenus", mr.menuApi.GetElTreeMenus)
	// without record
	base.GET("list", mr.menuApi.List)
}
