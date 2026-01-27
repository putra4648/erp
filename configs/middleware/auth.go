package middleware

import (
	"context"
	"putra4648/erp/configs/auth"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(verifier *oidc.IDTokenVerifier) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rawToken := c.Get("Authorization")
		if rawToken == "" || !strings.HasPrefix(rawToken, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token diperlukan"})
		}

		tokenString := strings.TrimPrefix(rawToken, "Bearer ")

		// 1. Verifikasi Signature & Expiration secara OFFLINE
		idToken, err := verifier.Verify(context.Background(), tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid: " + err.Error()})
		}

		// 2. Parse Claims untuk mengambil Roles
		var claims auth.KeycloakClaims
		if err := idToken.Claims(&claims); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal extract claims"})
		}

		// Simpan data ke context agar bisa dipakai di handler berikutnya
		c.Locals("user_id", claims.Subject)
		c.Locals("roles", claims.RealmAccess.Roles)
		return c.Next()
	}
}
