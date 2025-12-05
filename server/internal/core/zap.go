package core

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	coreZap "server/internal/core/zap"
	"server/internal/global"
	"server/internal/pkg"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := pkg.PathExists(global.TD27_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.TD27_CONFIG.Zap.Director)
		_ = os.Mkdir(global.TD27_CONFIG.Zap.Director, os.ModePerm)
	}

	levels := global.TD27_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := coreZap.NewZapCore(levels[i])
		cores = append(cores, core)
	}

	logger = zap.New(zapcore.NewTee(cores...))

	if global.TD27_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
