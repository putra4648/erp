package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func LoggerMiddleware(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Handle request
		err := c.Next()

		// Log request details
		duration := time.Since(start)

		status := c.Response().StatusCode()
		method := c.Method()
		url := c.OriginalURL()
		ip := c.IP()
		var permissionsStr string
		if p, ok := c.Locals("permissions").([]string); ok {
			permissionsStr = strings.Join(p, ",")
		}

		requiredPermission, _ := c.Locals("required_permission").(string)

		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("url", url),
			zap.String("ip", ip),
			zap.Duration("latency", duration),
			zap.String("user_agent", c.Get(fiber.HeaderUserAgent)),
			zap.String("permissions", permissionsStr),
			zap.String("required_permission", requiredPermission),
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
		}

		if status >= 500 {
			log.Error("HTTP request failed", fields...)
		} else if status >= 400 {
			log.Warn("HTTP request warning", fields...)
		} else {
			log.Info("HTTP request success", fields...)
		}

		return err
	}
}
