package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type BaseRouter struct{}

func (br *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")

	baseApi := api.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("login", baseApi.Login)
	}

	return baseRouter
}
