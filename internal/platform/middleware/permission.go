package middleware

import (
	"putra4648/erp/internal/platform/logger"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
)

// PermissionMiddleware checks if a user's roles grant them permission for the request.
// It assumes that a preceding authentication middleware has validated the user's JWT
// and placed their roles into c.Locals("roles") as a []string.
func PermissionMiddleware(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, ok := c.Locals("roles").([]string)
		if !ok {
			// No roles found, or not in the expected format.
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have permission to access this resource"})
		}

		path := c.Path()
		method := c.Method()
		isAllowed := false

		for _, role := range roles {
			allowed, err := e.Enforce(role, path, method)
			if err != nil {
				logger.Log.Errorf("Casbin enforce error: %v", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error during authorization"})
			}
			if allowed {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have permission to access this resource"})
	}
}
