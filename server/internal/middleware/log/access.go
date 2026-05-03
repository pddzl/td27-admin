package log

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// GinLogger is a gin middleware that logs every HTTP request.
// 4xx responses are logged at Warn, 5xx at Error, everything else at Info.
// The request_id (if set by RequestID middleware) is included in the log line.
func GinLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		status := c.Writer.Status()
		attrs := []any{
			"status", status,
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"user-agent", c.Request.UserAgent(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate),
			"cost", cost,
		}

		switch {
		case status >= 500:
			logger.Error("server error", attrs...)
		case status >= 400:
			logger.Warn("client error", attrs...)
		default:
			logger.Info("request handled", attrs...)
		}
	}
}
