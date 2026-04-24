package log

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(
			path,
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"user-agent", c.Request.UserAgent(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate),
			"cost", cost,
		)
	}
}
