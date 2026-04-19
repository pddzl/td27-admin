package sysTool

import (
	"github.com/gin-gonic/gin"

	apiSysTool "server/internal/api/sysTool"
	"server/internal/middleware"
)

type ServiceTokenRouter struct {
	serviceTokenApi *apiSysTool.ServiceTokenApi
}

func (r *ServiceTokenRouter) InitServiceTokenRouter(rg *gin.RouterGroup) {
	base := rg.Group("serviceToken")
	record := base.Use(middleware.OperationRecord())
	// record
	record.POST("create", r.serviceTokenApi.Create)
	record.GET("delete", r.serviceTokenApi.Delete)
	record.GET("update", r.serviceTokenApi.Update)
	record.GET("detail", r.serviceTokenApi.GetById)
	// without record
	base.POST("list", r.serviceTokenApi.List)
}
