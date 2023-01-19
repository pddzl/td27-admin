package core

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"server/core/internal"
	"server/global"
	"server/utils"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.TD27_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.TD27_CONFIG.Zap.Director)
		_ = os.Mkdir(global.TD27_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.TD27_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
