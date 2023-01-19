package internal

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"

	"server/global"
)

type lumberjackLogs struct{}

var LumberjackLogs = new(lumberjackLogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (l *lumberjackLogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename:   path.Join(global.TD27_CONFIG.Zap.Director, level+".log"),
		MaxSize:    global.TD27_CONFIG.RotateLogs.MaxSize,
		MaxBackups: global.TD27_CONFIG.RotateLogs.MaxBackups,
		MaxAge:     global.TD27_CONFIG.RotateLogs.MaxAge,
		Compress:   global.TD27_CONFIG.RotateLogs.Compress,
	}

	if global.TD27_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
