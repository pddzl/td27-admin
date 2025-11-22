package base

import (
	"github.com/gin-gonic/gin"
	
	"server/internal/api/base"
	"server/internal/middleware"
)

type CasbinRouter struct {
	casbinApi *base.CasbinApi
}

func NewCasbinRouter() *CasbinRouter {
	return &CasbinRouter{
		casbinApi: base.NewCasbinApi(),
	}
}

func (cr *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouter.POST("editCasbin", cr.casbinApi.EditCasbin)
	return casbinRouter
}
