package middlewares

import (
	"net/http"
	"strings"

	"to-do-api/utils"

	"github.com/gin-gonic/gin"
)

// JWT Authentication Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
