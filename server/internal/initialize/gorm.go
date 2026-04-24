package initialize

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"server/internal/global"
	"server/internal/model/sysManagement"
	modelMonitor "server/internal/model/sysMonitor"
	modelSysTool "server/internal/model/sysTool"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	logZap := global.TD27_CONFIG.Pgsql.LogZap
	if logZap {
		global.TD27_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}

func gormConfig() *gorm.Config {
	newLogger := logger.New(
		NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:                  logger.Warn,            // 日志级别
			IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                  // 禁用彩色打印
		},
	)
	config := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	return config
}

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	p := global.TD27_CONFIG.Pgsql

	if p.Dbname == "" {
		return nil
	}

	// PostgreSQL DSN 格式
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s",
		p.Host, p.Username, p.Password, p.Dbname, p.Port, p.Config)

	pgConfig := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false, // 启用 prepared statement 缓存以提高性能
	}

	if db, err := gorm.Open(postgres.New(pgConfig), gormConfig()); err != nil {
		global.TD27_LOG.Error("pgsql连接失败", "error", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)
		sqlDB.SetConnMaxIdleTime(30 * time.Minute)
		return db
	}
}

// RegisterTables 初始化数据库表
func RegisterTables(db *gorm.DB) {
	if !global.TD27_CONFIG.System.DisableAutoMigrate {
		err := db.AutoMigrate(
			// 权限 - 用户和角色（多对多）
			sysManagement.UserModel{},
			sysManagement.RoleModel{},
			sysManagement.UserRole{},
			// 菜单模型（数据来自 permission 表 type='menu'）
			sysManagement.MenuModel{},
			// API权限
			sysManagement.ApiModel{},
			// 统一权限模型
			sysManagement.PermissionModel{},
			sysManagement.RolePermissionModel{},
			// 部门（用于数据权限）
			sysManagement.DeptModel{},
			// 监控
			modelMonitor.OperationLogModel{},
			// file
			modelSysTool.FileModel{},
			// crontab
			modelSysTool.CronModel{},
			// 服务令牌
			modelSysTool.ServiceToken{},
			modelSysTool.TokenPermission{},
			// 缓存表
			modelSysTool.CacheModel{},
			// 字典
			sysManagement.DictModel{},
			sysManagement.DictDetailModel{},
			// 按钮权限
			sysManagement.ButtonModel{},
		)

		if err != nil {
			// 忽略"已存在"错误，这是正常的当init.sql已经创建了约束
			if isAlreadyExistsError(err) {
				global.TD27_LOG.Info("AutoMigrate: some constraints already exist (from init.sql), continuing...")
			} else if isNotExistsError(err) {
				// 忽略"不存在"错误，可能是GORM尝试删除不存在的约束
				global.TD27_LOG.Info("AutoMigrate: constraint does not exist, continuing...")
			} else {
				global.TD27_LOG.Error("register table failed", "error", err)
				os.Exit(0)
			}
		}
	}
	global.TD27_LOG.Info("register table success")
}

// isAlreadyExistsError 检查错误是否是"已存在"错误
func isAlreadyExistsError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	// MySQL: already exists
	// PostgreSQL: already exists
	return strings.Contains(errStr, "already exists") ||
		strings.Contains(errStr, "Duplicate key name") ||
		strings.Contains(errStr, "42P07") || // PostgreSQL: duplicate_table
		strings.Contains(errStr, "42710") // PostgreSQL: duplicate_object
}

// isNotExistsError 检查错误是否是"不存在"错误（尝试删除不存在的约束）
func isNotExistsError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "does not exist") ||
		strings.Contains(errStr, "42704") // PostgreSQL: undefined_object
}
