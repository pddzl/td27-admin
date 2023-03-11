package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type JwtRouter struct{}

func (jr *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	jwtRouter := Router.Group("jwt")
	jwtApi := api.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.POST("joinInBlacklist", jwtApi.JoinInBlacklist)
	}
	return jwtRouter
}
