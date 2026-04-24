package core

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"

	"server/internal/global"
	"server/internal/pkg"
)

func Zap() *slog.Logger {
	dir := global.TD27_CONFIG.Zap.Director
	if ok, _ := pkg.PathExists(dir); !ok {
		fmt.Printf("create %v directory\n", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}

	level := parseLevel(global.TD27_CONFIG.Zap.Level)

	opts := &slog.HandlerOptions{
		Level: level,
	}
	if global.TD27_CONFIG.Zap.ShowLine {
		opts.AddSource = true
	}

	var handler slog.Handler
	if global.TD27_CONFIG.Zap.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	fileWriter := &lumberjack.Logger{
		Filename:   path.Join(dir, "app.log"),
		MaxSize:    global.TD27_CONFIG.RotateLogs.MaxSize,
		MaxBackups: global.TD27_CONFIG.RotateLogs.MaxBackups,
		MaxAge:     global.TD27_CONFIG.RotateLogs.MaxAge,
		Compress:   global.TD27_CONFIG.RotateLogs.Compress,
	}

	var fileHandler slog.Handler
	if global.TD27_CONFIG.Zap.Format == "json" {
		fileHandler = slog.NewJSONHandler(fileWriter, opts)
	} else {
		fileHandler = slog.NewTextHandler(fileWriter, opts)
	}

	mh := &multiHandler{handlers: []slog.Handler{handler, fileHandler}}
	logger := slog.New(mh)
	slog.SetDefault(logger)
	return logger
}

func parseLevel(lvl string) slog.Level {
	switch lvl {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

type multiHandler struct {
	handlers []slog.Handler
}

func (m *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *multiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		_ = h.Handle(ctx, r)
	}
	return nil
}

func (m *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithAttrs(attrs)
	}
	return &multiHandler{handlers: newHandlers}
}

func (m *multiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &multiHandler{handlers: newHandlers}
}
