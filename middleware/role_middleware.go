package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

// Middleware khusus pengecekan role
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles := c.MustGet("roles").([]string)

		if slices.Contains(roles, requiredRole) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Anda tidak memiliki akses " + requiredRole})
	}
}
