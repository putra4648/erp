package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// RequirePermission checks if a user has the required permission for the request.
// It assumes that a preceding authentication middleware has validated the user's JWT
// and placed their permissions into c.Locals("roles") as a []string.
func RequirePermission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		permissions, ok := c.Locals("roles").([]string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have permission to access this resource"})
		}

		for _, p := range permissions {
			if p == requiredPermission {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have permission to access this resource"})
	}
}
