package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type ApiRouter struct{}

func (u *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiWithoutRouter := Router.Group("api")

	apiApi := api.ApiGroupApp.SystemApiGroup.ApiApi
	{
		apiRouter.POST("addApi", apiApi.AddApi)
		apiRouter.POST("deleteApi", apiApi.DeleteApi)
		apiRouter.POST("editApi", apiApi.EditApi)
		apiRouter.POST("getElTreeApis", apiApi.GetElTreeApis)
	}
	{
		apiWithoutRouter.POST("getApis", apiApi.GetApis)
	}
}
