package sysManagement

import (
	"github.com/gin-gonic/gin"

	apiSysManagement "server/internal/api/sysManagement"
	"server/internal/middleware"
)

type ButtonRouter struct{}

func (r *ButtonRouter) InitButtonRouter(Router *gin.RouterGroup) {
	buttonRouter := Router.Group("button").Use(middleware.OperationRecord())
	buttonApi := apiSysManagement.NewButtonApi()
	{
		buttonRouter.POST("create", buttonApi.Create)
		buttonRouter.POST("delete", buttonApi.Delete)
		buttonRouter.POST("update", buttonApi.Update)
		buttonRouter.POST("list", buttonApi.List)
		buttonRouter.GET("page", buttonApi.GetPageButtons)
		buttonRouter.POST("check", buttonApi.CheckPermission)
		buttonRouter.POST("batchCheck", buttonApi.BatchCheckPermission)
		buttonRouter.GET("user", buttonApi.GetUserButtons)
	}
}
