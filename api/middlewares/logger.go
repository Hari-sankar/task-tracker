package middleware

import (
	"task-tracker/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		logger.Info("Request started",
			zap.String("path", path),
			zap.String("method", c.Request.Method),
			zap.String("client_ip", c.ClientIP()),
		)

		c.Next()

		logger.Info("Request completed",
			zap.String("path", path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}
