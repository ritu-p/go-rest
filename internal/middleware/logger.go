package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()
		// keep minimal to avoid importing heavy logger libs here
		c.Writer.Header().Set("X-Processed-Time", latency.String())
		// you can hook your logger here
		_ = status
	}
}
