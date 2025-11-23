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

func (cr *CasbinRouter) InitCasbinRouter(rg *gin.RouterGroup) {
	baseG := rg.Group("casbin")
	record := baseG.Use(middleware.OperationRecord())
	record.POST("editCasbin", cr.casbinApi.EditCasbin)
}
