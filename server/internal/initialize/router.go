package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "server/docs"
	"server/internal/global"
	"server/internal/middleware"
	"server/internal/middleware/log"
	"server/internal/router"
)

func Routers() *gin.Engine {
	if global.TD27_CONFIG.System.Env == "prod" {
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
	PublicGroup := Router.Group(global.TD27_CONFIG.Router.Prefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	// 路由组
	// -> 基础
	baseRouter := router.NewBaseRouterGroup()
	{
		baseRouter.InitLogRegRouter(PublicGroup) // 登录相关
	}

	// -> 鉴权管理
	authorityRouter := router.NewAuthorityRouterGroup()
	// -> 文件管理
	fileMRouter := router.NewFileMRouterGroup()
	// -> 系统监控
	monitorRouter := router.NewMonitorRouterGroup()
	// -> 系统工具
	sysToolRouter := router.NewSysToolRouterGroup()
	// -> system settings
	sysSetRouter := router.NewSysSetRouterGroup()

	// 需要认证的路由
	PrivateGroup := Router.Group(global.TD27_CONFIG.Router.Prefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		// 基础
		baseRouter.InitCasbinRouter(PrivateGroup)
		// 鉴权管理
		authorityRouter.InitUserRouter(PrivateGroup) // 用户
		authorityRouter.InitRoleRouter(PrivateGroup) // 角色
		authorityRouter.InitMenuRouter(PrivateGroup) // 菜单
		authorityRouter.InitApiRouter(PrivateGroup)  // 接口
		// 系统监控
		monitorRouter.InitOperationLogRouter(PrivateGroup) // 操作日志
		// 文件管理
		fileMRouter.InitFileRouter(PrivateGroup)
		// 系统工具
		sysToolRouter.InitCronRouter(PrivateGroup)
		// system settings
		sysSetRouter.InitDictRouter(PrivateGroup)
		sysSetRouter.InitDictDetailRouter(PrivateGroup)
	}

	global.TD27_LOG.Info("router register success")

	return Router
}
