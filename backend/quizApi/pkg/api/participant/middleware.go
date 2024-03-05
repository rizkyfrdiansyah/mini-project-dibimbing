package participant

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
		if !ok || userRole != "participant" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: User is not a participant"})
			c.Abort()
			return
		}

		// Call the next middleware/handler in chain
		c.Next()
	}
}
