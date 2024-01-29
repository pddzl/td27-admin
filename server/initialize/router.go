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

	// 公共路由组 不需要鉴权
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	// 路由组
	// -> 基础
	baseRouter := router.RouterGroupApp.Base
	{
		baseRouter.InitLogRegRouter(PublicGroup) // 登录相关
	}

	// -> 鉴权管理
	authorityRouter := router.RouterGroupApp.Authority
	// -> 文件管理
	fileMRouter := router.RouterGroupApp.FileM
	// -> 系统监控
	monitorRouter := router.RouterGroupApp.Monitor

	// 需要认证的路由
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		// 基础
		baseRouter.InitCasbinRouter(PrivateGroup)
		baseRouter.InitJwtRouter(PrivateGroup)
		// 鉴权管理
		authorityRouter.InitUserRouter(PrivateGroup) // 用户
		authorityRouter.InitRoleRouter(PrivateGroup) // 角色
		authorityRouter.InitMenuRouter(PrivateGroup) // 菜单
		authorityRouter.InitApiRouter(PrivateGroup)  // 接口
		// 系统监控
		monitorRouter.InitOperationLogRouter(PrivateGroup) // 操作日志
		// 文件管理
		fileMRouter.InitFileRouter(PrivateGroup)
	}

	global.TD27_LOG.Info("router register success")
	return Router
}
