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
	"log/slog"
)

func Routers() *gin.Engine {
	if global.TD27_CONFIG.System.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(log.GinLogger(global.TD27_LOG), log.GinRecovery(global.TD27_LOG, global.TD27_CONFIG.System.Stack))

	// 跨域，如需跨域可以打开下面的注释
	// slog.Info("use middleware cors")
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public group
	publicGroup := r.Group(global.TD27_CONFIG.System.RouterPrefix)
	publicGroup.GET("/health", func(c *gin.Context) { c.JSON(200, "ok") })

	// Private group
	privateGroup := r.Group(global.TD27_CONFIG.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth(), middleware.CasbinHandler())

	// register all router
	router.RegisterAllModuleRouter()

	// Automatically load ALL router modules
	for _, m := range router.GetAllModules() {
		m.InitPublic(publicGroup)
		m.InitPrivate(privateGroup)
	}

	slog.Info("router register success")

	return r
}
