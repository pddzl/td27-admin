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

	opts := &slog.HandlerOptions{
		Level: level,
	}
	if global.TD27_CONFIG.Logger.ShowLine {
		opts.AddSource = true
	}

	newHandler := func(w io.Writer) slog.Handler {
		if global.TD27_CONFIG.Logger.Format == "json" {
			return slog.NewJSONHandler(w, opts)
		}
		return slog.NewTextHandler(w, opts)
	}

	var handlers []slog.Handler
	if global.TD27_CONFIG.Logger.LogInConsole {
		handlers = append(handlers, newHandler(os.Stdout))
	}

	handlers = append(handlers, newHandler(&lumberjack.Logger{
		Filename:   path.Join(dir, "app.log"),
		MaxSize:    global.TD27_CONFIG.RotateLogs.MaxSize,
		MaxBackups: global.TD27_CONFIG.RotateLogs.MaxBackups,
		MaxAge:     global.TD27_CONFIG.RotateLogs.MaxAge,
		Compress:   global.TD27_CONFIG.RotateLogs.Compress,
	}))

	// Build static attrs injected on every log record
	staticAttrs := buildStaticAttrs()

	mh := &multiHandler{handlers: handlers}
	sh := &staticAttrHandler{handler: mh, attrs: staticAttrs}
	logger := slog.New(sh)
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

// staticAttrHandler wraps a handler to prepend fixed attributes to every record.
type staticAttrHandler struct {
	handler slog.Handler
	attrs   []slog.Attr
}

func (h *staticAttrHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *staticAttrHandler) Handle(ctx context.Context, r slog.Record) error {
	// Clone to avoid mutating the original; add static attrs before delegating
	r2 := r.Clone()
	r2.AddAttrs(h.attrs...)
	return h.handler.Handle(ctx, r2)
}

func (h *staticAttrHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &staticAttrHandler{handler: h.handler.WithAttrs(attrs), attrs: h.attrs}
}

func (h *staticAttrHandler) WithGroup(name string) slog.Handler {
	return &staticAttrHandler{handler: h.handler.WithGroup(name), attrs: h.attrs}
}

// multiHandler fans out log records to multiple handlers (console + file).
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
	var firstErr error
	for i, h := range m.handlers {
		if err := h.Handle(ctx, r); err != nil {
			if firstErr == nil {
				firstErr = err
			}
			// Fallback: log handler failure to stdout handler
			if i > 0 && len(m.handlers) > 0 {
				_ = m.handlers[0].Handle(ctx, slog.Record{
					Time:    r.Time,
					Level:   slog.LevelError,
					Message: fmt.Sprintf("log handler %d failed: %v", i, err),
				})
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
	return &multiHandler{handlers: newHandlers}
}

func (m *multiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		newHandlers[i] = h.WithGroup(name)
	}
	return &multiHandler{handlers: newHandlers}
}
