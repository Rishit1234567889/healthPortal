package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware logs request details
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		// Process request
		ctx.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get client IP
		clientIP := ctx.ClientIP()

		// Get response status
		status := ctx.Writer.Status()

		// Get request method
		method := ctx.Request.Method

		// Create query string
		query := ""
		if raw != "" {
			query = "?" + raw
		}

		// Log the request details
		logger.Info("HTTP Request",
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", clientIP),
			zap.Duration("latency", latency),
		)
	}
}
