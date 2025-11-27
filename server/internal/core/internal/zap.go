package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"server/internal/global"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	entity := &ZapCore{level: level}
	syncer := LumberjackLogs.GetWriteSyncer(level.String())
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(global.TD27_CONFIG.Zap.Encoder(), syncer, levelEnabler)
	return entity
}
