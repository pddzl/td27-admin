package core

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"server/configs"
	"server/internal/global"
)

// Viper 初始化配置
// 配置文件优先级：环境变量 TD27_CONFIG > 命令行参数 -c > 默认 configs/config.yaml
func Viper() *viper.Viper {
	// 从环境变量获取配置文件路径
	config := os.Getenv("TD27_CONFIG")
	if config == "" {
		// 尝试从命令行参数获取
		for i, arg := range os.Args {
			if arg == "-c" && i+1 < len(os.Args) {
				config = os.Args[i+1]
				break
			}
		}
	}
	if config == "" {
		config = "configs/config.yaml"
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error configs file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("configs file changed:", e.Name)
		if err = v.Unmarshal(&global.TD27_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.TD27_CONFIG); err != nil {
		panic(fmt.Errorf("Failed to unmarshal config: %s \n", err))
	}

	// 验证关键配置
	if err = validateConfig(&global.TD27_CONFIG); err != nil {
		panic(fmt.Errorf("Config validation failed: %s \n", err))
	}

	return v
}

// validateConfig 验证关键配置项
func validateConfig(cfg *configs.Server) error {
	// 验证 JWT 配置
	if cfg.JWT.SigningKey == "" {
		return fmt.Errorf("JWT signing key is required")
	}
	if cfg.JWT.ExpiresTime <= 0 {
		return fmt.Errorf("JWT expires time must be greater than 0")
	}

	// 验证 MySQL 配置
	if cfg.Pgsql.Host == "" {
		return fmt.Errorf("PgSQL host is required")
	}
	if cfg.Pgsql.Port == "" {
		return fmt.Errorf("PgSQL port is required")
	}
	if cfg.Pgsql.Dbname == "" {
		return fmt.Errorf("PgSQL database name is required")
	}
	if cfg.Pgsql.Username == "" {
		return fmt.Errorf("PgSQL username is required")
	}

	// 验证系统配置
	if cfg.System.Port <= 0 {
		return fmt.Errorf("system port must be greater than 0")
	}

	return nil
}
