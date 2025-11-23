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

func (br *LogRegRouter) InitLogRegRouter(rg *gin.RouterGroup) {
	baseG := rg.Group("logReg")
	baseG.POST("captcha", br.logRegApi.Captcha)
	baseG.POST("login", br.logRegApi.Login)
	baseG.POST("logout", br.logRegApi.LogOut)
}
