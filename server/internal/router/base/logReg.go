package base

import (
	"github.com/gin-gonic/gin"
	"server/internal/api/base"
)

type LogRegRouter struct {
	logRegApi *base.LogRegApi
}

func NewLogRegRouter() *LogRegRouter {
	return &LogRegRouter{
		logRegApi: base.NewLogRegApi(),
	}
}

func (br *LogRegRouter) InitLogRegRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	logRegRouter := Router.Group("logReg")
	logRegRouter.POST("captcha", br.logRegApi.Captcha)
	logRegRouter.POST("login", br.logRegApi.Login)
	logRegRouter.POST("logout", br.logRegApi.LogOut)

	return logRegRouter
}
