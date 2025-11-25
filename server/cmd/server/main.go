package main

import (
	"fmt"

	"go.uber.org/zap"

	"server/internal/core"
	"server/internal/global"
	"server/internal/initialize"
)

func main() {
	// Load configuration first (everything depends on it)
	global.TD27_VP = core.Viper()

	// Setup logger
	global.TD27_LOG = core.Zap()
	zap.ReplaceGlobals(global.TD27_LOG)

	// Initialize MySQL
	global.TD27_DB = initialize.Gorm()
	if global.TD27_DB == nil {
		global.TD27_LOG.Fatal("mysql connection failed")
	}
	db, _ := global.TD27_DB.DB()
	defer db.Close()

	global.TD27_REDIS = initialize.Redis()
	if global.TD27_REDIS == nil {
		global.TD27_LOG.Fatal("redis connection failed")
	}

	// Initialize Cron AFTER DB/Redis ready
	global.TD27_CRON = initialize.InitCron()
	initialize.CheckCron()

	// Auto migrate tables AFTER DB is initialized
	initialize.RegisterTables(global.TD27_DB)

	// Build router
	router := initialize.Routers()

	// Run HTTP server
	addr := fmt.Sprintf("%s:%d", global.TD27_CONFIG.System.Host, global.TD27_CONFIG.System.Port)
	core.RunServer(addr, router)
}
