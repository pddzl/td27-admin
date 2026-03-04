package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/middleware"
)

type CasbinRouter struct {
	casbinApi *sysManagement.CasbinApi
}

func NewCasbinRouter() *CasbinRouter {
	return &CasbinRouter{
		casbinApi: sysManagement.NewCasbinApi(),
	}
}

func (cr *CasbinRouter) InitCasbinRouter(rg *gin.RouterGroup) {
	baseG := rg.Group("casbin")
	record := baseG.Use(middleware.OperationRecord())
	record.POST("update", cr.casbinApi.Update)
}
