package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/internal/middleware"
	"server/internal/router"

	_ "server/docs"
	"server/internal/global"
	"server/internal/middleware/log"
)

func Routers() *gin.Engine {
	if global.TD27_CONFIG.System.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(log.GinLogger(), log.GinRecovery(global.TD27_CONFIG.System.Stack))

	// 跨域，如需跨域可以打开下面的注释
	// global.GVA_LOG.Info("use middleware cors")
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public group
	publicGroup := r.Group(global.TD27_CONFIG.Router.Prefix)
	publicGroup.GET("/health", func(c *gin.Context) { c.JSON(200, "ok") })

	// Private group
	privateGroup := r.Group(global.TD27_CONFIG.Router.Prefix)
	privateGroup.Use(middleware.JWTAuth(), middleware.CasbinHandler())

	// Automatically load ALL router modules
	for _, m := range router.GetAllModules() {
		m.InitPublic(publicGroup)
		m.InitPrivate(privateGroup)
	}

	global.TD27_LOG.Info("router register success")

	return r
}
