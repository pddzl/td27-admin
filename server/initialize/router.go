package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server/docs"
	"server/global"
	"server/middleware"
	"server/middleware/log"
	"server/router"
)

func Routers() *gin.Engine {
	if global.TD27_CONFIG.System.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	Router := gin.New()
	Router.Use(log.GinLogger(), log.GinRecovery(global.TD27_CONFIG.System.Stack))

	// 跨域，如需跨域可以打开下面的注释
	// global.GVA_LOG.Info("use middleware cors")
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	global.TD27_LOG.Info("register swagger handler")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由组
	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	// 需要认证的路由
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitRoleRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitApiRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitOperationRecordRouter(PrivateGroup)
	}

	global.TD27_LOG.Info("router register success")
	return Router
}
