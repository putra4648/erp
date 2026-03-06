package middleware

import (
	"putra4648/erp/configs/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func AuthMiddleware(a *auth.Authenticator, log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Required authorization header is missing",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Authorization header format must be Bearer {token}",
			})
		}

		tokenString := parts[1]
		accessToken, err := a.VerifyToken(c.Context(), tokenString)
		if err != nil {
			log.Error("JWT validation failed", zap.Error(err), zap.String("path", c.Path()))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Failed to validate JWT.",
			})
		}

		var profile map[string]interface{}
		if err := accessToken.Claims(&profile); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Extract permissions for RequirePermission
		var userPermissions []string
		if permissions, ok := profile["permissions"].([]interface{}); ok {
			for _, p := range permissions {
				if s, ok := p.(string); ok {
					userPermissions = append(userPermissions, s)
				}
			}
		}
		c.Locals("permissions", userPermissions)

		// Also add scope if useful
		var userScopes []string
		if scope, ok := profile["scope"].(string); ok {
			userScopes = append(userPermissions, strings.Split(scope, " ")...)
		}
		c.Locals("scopes", userScopes)

		return c.Next()
	}
}
