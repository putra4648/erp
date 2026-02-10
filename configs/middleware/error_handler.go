package middleware

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RecoverMiddleware handles panics and logs them using Zap
func RecoverMiddleware(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Error("Panic recovered",
					zap.Any("panic", r),
					zap.String("stack", string(debug.Stack())),
					zap.String("url", c.OriginalURL()),
					zap.String("method", c.Method()),
				)
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Internal Server Error",
				})
			}
		}()
		return c.Next()
	}
}

// GlobalErrorHandler logs any error returned from handlers
func GlobalErrorHandler(log *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// Log the error
		log.Error("Application error",
			zap.Error(err),
			zap.Int("status", code),
			zap.String("method", c.Method()),
			zap.String("url", c.OriginalURL()),
		)

		// Return status code with error message
		return c.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}
