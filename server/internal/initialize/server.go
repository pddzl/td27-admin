package initialize

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"server/internal/global"
	"syscall"
	"time"
)

// RunServer starts the HTTP server and handles graceful shutdown
func RunServer(addr string, handler http.Handler) {
	srv := &http.Server{
		Addr:           addr,
		Handler:        handler,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Start server in a goroutine
	go func() {
		//global.TD27_LOG.Info("server listening", zap.String("addr", addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.TD27_LOG.Error("http server failed ", zap.Any("err", err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.TD27_LOG.Info("shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.TD27_LOG.Error("server forced to shutdown", zap.Any("err", err))
	}

	global.TD27_LOG.Info("server exiting")
}
