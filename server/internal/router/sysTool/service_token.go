package sysTool

import (
	"github.com/gin-gonic/gin"

	apiSysTool "server/internal/api/sysTool"
	"server/internal/middleware"
)

type ServiceTokenRouter struct{}

func (s *ServiceTokenRouter) InitServiceTokenRouter(Router *gin.RouterGroup) {
	serviceTokenRouter := Router.Group("serviceToken").Use(middleware.OperationRecord())
	serviceTokenApi := apiSysTool.NewServiceTokenApi()
	{
		serviceTokenRouter.POST("create", serviceTokenApi.Create)
		serviceTokenRouter.POST("delete", serviceTokenApi.Delete)
		serviceTokenRouter.POST("update", serviceTokenApi.Update)
		serviceTokenRouter.POST("detail", serviceTokenApi.GetById)
		serviceTokenRouter.POST("list", serviceTokenApi.List)
	}
}
