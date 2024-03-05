package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User not found"})
			c.Abort()
			return
		}

		userRole, ok := user.(string)
		if !ok || userRole != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User is not an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
