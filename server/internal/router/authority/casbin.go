package authority

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/authority"
	"server/internal/middleware"
)

type CasbinRouter struct {
	casbinApi *authority.CasbinApi
}

func NewCasbinRouter() *CasbinRouter {
	return &CasbinRouter{
		casbinApi: authority.NewCasbinApi(),
	}
}

func (cr *CasbinRouter) InitCasbinRouter(rg *gin.RouterGroup) {
	baseG := rg.Group("casbin")
	record := baseG.Use(middleware.OperationRecord())
	record.POST("editCasbin", cr.casbinApi.EditCasbin)
}
