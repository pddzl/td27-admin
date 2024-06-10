package base

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type LogRegRouter struct{}

func (br *LogRegRouter) InitLogRegRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	logRegRouter := Router.Group("logReg")

	logRegApi := api.ApiGroupApp.Base.LogRegApi
	{
		logRegRouter.POST("captcha", logRegApi.Captcha)
		logRegRouter.POST("login", logRegApi.Login)
		logRegRouter.POST("logout", logRegApi.LogOut)
	}

	return logRegRouter
}
