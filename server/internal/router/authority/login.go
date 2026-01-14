package authority

import (
	"server/internal/api/authority"

	"github.com/gin-gonic/gin"
)

type LogRegRouter struct {
	logRegApi *authority.LogRegApi
}

func NewLogRegRouter() *LogRegRouter {
	return &LogRegRouter{
		logRegApi: authority.NewLogRegApi(),
	}
}

func (br *LogRegRouter) InitLogRegRouter(rg *gin.RouterGroup) {
	baseG := rg.Group("logReg")
	baseG.POST("captcha", br.logRegApi.Captcha)
	baseG.POST("login", br.logRegApi.Login)
	baseG.POST("logout", br.logRegApi.LogOut)
}
