package base

import (
	"github.com/gin-gonic/gin"
	"server/internal/api"
	"server/internal/middleware"
)

type CasbinRouter struct{}

func (cr *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinApi := api.ApiGroupApp.Base.CasbinApi
	{
		casbinRouter.POST("editCasbin", casbinApi.EditCasbin)
	}
	return casbinRouter
}
