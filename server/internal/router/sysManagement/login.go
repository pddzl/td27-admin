package sysManagement

import (
	"github.com/gin-gonic/gin"

	"server/internal/api/sysManagement"
	"server/internal/global"
	"server/internal/middleware"
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

	// 登录/验证码接口单独限流，防暴力破解
	authG := baseG.Group("")
	if cfg := global.TD27_CONFIG.RateLimit; cfg.Enabled {
		loginRate, loginBurst := cfg.LoginRate, cfg.LoginBurst
		if loginRate <= 0 {
			loginRate = 0.5
		}
		if loginBurst <= 0 {
			loginBurst = 5
		}
		authG.Use(middleware.NewIPRateLimiter(loginRate, loginBurst).Middleware())
	}

	authG.POST("captcha", r.logRegApi.Captcha)
	authG.POST("login", r.logRegApi.Login)
	baseG.POST("logout", r.logRegApi.LogOut)
}
