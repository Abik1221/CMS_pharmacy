package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireRoles checks if the user has one of the allowed roles
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context (set after AuthMiddleware)
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		// Check if the user's role is allowed
		roleValid := false
		for _, role := range allowedRoles {
			if role == userRole {
				roleValid = true
				break
			}
		}

		if !roleValid {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
