package core

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"

	"server/internal/global"
	"server/internal/pkg"
)

func Logger() *slog.Logger {
	dir := global.TD27_CONFIG.Logger.Director
	if ok, _ := pkg.PathExists(dir); !ok {
		fmt.Printf("create %v directory\n", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}

	level := parseLevel(global.TD27_CONFIG.Logger.Level)

	consoleOpts := &slog.HandlerOptions{
		Level: level,
	}
	fileOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug, // lowest level so the inner handler doesn't reject records before levelFilter
	}
	if global.TD27_CONFIG.Logger.ShowLine {
		consoleOpts.AddSource = true
		fileOpts.AddSource = true
	}

	newConsoleHandler := func(w io.Writer) slog.Handler {
		if global.TD27_CONFIG.Logger.Format == "json" {
			return slog.NewJSONHandler(w, consoleOpts)
		}
		return slog.NewTextHandler(w, consoleOpts)
	}

	newFileHandler := func(w io.Writer) slog.Handler {
		if global.TD27_CONFIG.Logger.Format == "json" {
			return slog.NewJSONHandler(w, fileOpts)
		}
		return slog.NewTextHandler(w, fileOpts)
	}

	var handlers []slog.Handler
	if global.TD27_CONFIG.Logger.LogInConsole {
		handlers = append(handlers, newConsoleHandler(os.Stdout))
	}

	for _, lvl := range []struct {
		name  string
		level slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"error", slog.LevelError},
	} {
		handlers = append(handlers, &levelFilter{
			handler: newFileHandler(&lumberjack.Logger{
				Filename:   path.Join(dir, lvl.name+".log"),
				MaxSize:    global.TD27_CONFIG.RotateLogs.MaxSize,
				MaxBackups: global.TD27_CONFIG.RotateLogs.MaxBackups,
				MaxAge:     global.TD27_CONFIG.RotateLogs.MaxAge,
				Compress:   global.TD27_CONFIG.RotateLogs.Compress,
			}),
			level: lvl.level,
		})
	}

	// Build static attrs injected on every log record
	staticAttrs := buildStaticAttrs()

	mh := &multiHandler{handlers: handlers, attrs: staticAttrs}
	logger := slog.New(mh)
	slog.SetDefault(logger)
	return logger
}

func buildStaticAttrs() []slog.Attr {
	var attrs []slog.Attr
	if svc := global.TD27_CONFIG.Logger.Service; svc != "" {
		attrs = append(attrs, slog.String("service", svc))
	}
	if env := global.TD27_CONFIG.System.Env; env != "" {
		attrs = append(attrs, slog.String("env", env))
	}
	return attrs
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
		return slog.LevelInfo // safe default: don't flood production
	}
}

type levelFilter struct {
	handler slog.Handler
	level   slog.Level
}

func (f *levelFilter) Enabled(ctx context.Context, l slog.Level) bool {
	return l == f.level && f.handler.Enabled(ctx, l)
}

func (f *levelFilter) Handle(ctx context.Context, r slog.Record) error {
	return f.handler.Handle(ctx, r)
}

func (f *levelFilter) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &levelFilter{handler: f.handler.WithAttrs(attrs), level: f.level}
}

func (f *levelFilter) WithGroup(name string) slog.Handler {
	return &levelFilter{handler: f.handler.WithGroup(name), level: f.level}
}

// multiHandler fans out log records to multiple handlers (console + file).
type multiHandler struct {
	handlers []slog.Handler
	attrs    []slog.Attr // static attrs injected on every record before dispatch
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
	// Clone once and inject static attrs before dispatching to children
	if len(m.attrs) > 0 {
		r2 := r.Clone()
		r2.AddAttrs(m.attrs...)
		r = r2
	}

	var firstErr error
	for i, h := range m.handlers {
		if !h.Enabled(ctx, r.Level) {
			continue
		}
		if err := h.Handle(ctx, r); err != nil {
			if firstErr == nil {
				firstErr = err
			}
			// Fallback: log handler failure to first (console) handler
			if i > 0 && len(m.handlers) > 0 {
				_ = m.handlers[0].Handle(ctx, r)
			}
		}
	}
	return firstErr
}

func (m *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithAttrs(attrs)
	}
	return &multiHandler{handlers: newHandlers, attrs: m.attrs}
}

func (m *multiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &multiHandler{handlers: newHandlers, attrs: m.attrs}
}
