package main

import (
	"go.uber.org/zap"
	"os"

	"server/core"
	"server/global"
	"server/initialize"
)

func main() {
	global.TD27_VP = core.Viper() // 初始化viper
	global.TD27_LOG = core.Zap()  // 初始化zap日志
	zap.ReplaceGlobals(global.TD27_LOG)
	global.TD27_DB = initialize.Gorm()       // gorm连接数据库
	initialize.Redis()                       // 初始化redis
	global.TD27_CRON = initialize.InitCron() // 初始化cron
	initialize.CheckCron()                   // start cron entry, if exists
	if global.TD27_DB == nil {
		global.TD27_LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.TD27_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.TD27_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}
