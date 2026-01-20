package middleware

import (
	"context"
	"net/http"
	"putra4648/erp/types"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(verifier *oidc.IDTokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := c.GetHeader("Authorization")
		if rawToken == "" || !strings.HasPrefix(rawToken, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token diperlukan"})
			return
		}

		tokenString := strings.TrimPrefix(rawToken, "Bearer ")

		// 1. Verifikasi Signature & Expiration secara OFFLINE
		idToken, err := verifier.Verify(context.Background(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid: " + err.Error()})
			return
		}

		// 2. Parse Claims untuk mengambil Roles
		var claims types.KeycloakClaims
		if err := idToken.Claims(&claims); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Gagal extract claims"})
			return
		}

		// Simpan data ke context agar bisa dipakai di handler berikutnya
		c.Set("user_id", claims.Subject)
		c.Set("roles", claims.RealmAccess.Roles)
		c.Next()
	}
}
