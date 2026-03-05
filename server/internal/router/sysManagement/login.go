package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
)

type LogRegRouter struct {
	logRegApi *sysManagement.LogRegApi
}

func NewLogRegRouter() *LogRegRouter {
	return &LogRegRouter{
		logRegApi: sysManagement.NewLogRegApi(),
	}
}

func (r *LogRegRouter) InitLogRegRouter(rg *gin.RouterGroup) {
	baseG := rg.Group("")
	baseG.POST("captcha", r.logRegApi.Captcha)
	baseG.POST("login", r.logRegApi.Login)
	baseG.POST("logout", r.logRegApi.LogOut)
}
