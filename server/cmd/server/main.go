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

	// Initialize Database
	global.TD27_DB = initialize.Gorm()
	if global.TD27_DB == nil {
		global.TD27_LOG.Fatal("database connection failed")
	}
	db, _ := global.TD27_DB.DB()
	defer db.Close()

	// Auto migrate tables AFTER DB is initialized
	initialize.RegisterTables(global.TD27_DB)

	// Initialize Cron AFTER DB ready
	global.TD27_CRON = initialize.InitCron()
	initialize.CheckCron()

	// Build router
	router := initialize.Routers()

	// Run HTTP server
	addr := fmt.Sprintf("%s:%d", global.TD27_CONFIG.System.Host, global.TD27_CONFIG.System.Port)
	initialize.RunServer(addr, router)
}
