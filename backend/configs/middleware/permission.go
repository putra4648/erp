package middleware

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// RequirePermission checks if a user has the required permission for the request.
// It assumes that a preceding authentication middleware has validated the user's JWT
// and placed their permissions into c.Locals("roles") as a []string.
func RequirePermission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("required_permission", requiredPermission)
		permissions, ok := c.Locals("permissions").([]string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have permission to access this resource"})
		}

		requiredPermissions := strings.SplitSeq(requiredPermission, ",")

		for required := range requiredPermissions {
			if slices.Contains(permissions, required) {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You don't have permission to access this resource"})
	}
}
