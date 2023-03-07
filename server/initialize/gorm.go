package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"

	"server/global"
	"server/model/system"
)

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	m := global.TD27_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// RegisterTables 初始化数据库表
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.RoleModel{},
		system.UserModel{},
		system.MenuModel{},
		system.ApiModel{},
		system.JwtBlacklist{},
	)

	if err != nil {
		global.TD27_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.TD27_LOG.Info("register table success")
}
