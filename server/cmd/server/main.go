package main

import (
	"fmt"
	
	"go.uber.org/zap"

	"server/internal/core"
	"server/internal/global"
	"server/internal/initialize"
)

func main() {
	// Initialization
	global.TD27_VP = core.Viper()
	global.TD27_LOG = core.Zap()
	zap.ReplaceGlobals(global.TD27_LOG)

	global.TD27_DB = initialize.Gorm()
	initialize.Redis()
	global.TD27_CRON = initialize.InitCron()
	initialize.CheckCron()

	if global.TD27_DB == nil {
		global.TD27_LOG.Fatal("mysql connection failed")
	}

	initialize.RegisterTables(global.TD27_DB)
	db, _ := global.TD27_DB.DB()
	defer db.Close()

	// Build router
	router := initialize.Routers()
	addr := fmt.Sprintf("%s:%d", global.TD27_CONFIG.System.Host, global.TD27_CONFIG.System.Port)

	// Run HTTP server
	core.RunServer(addr, router)
}
