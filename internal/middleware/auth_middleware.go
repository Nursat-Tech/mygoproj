package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nursat/myproj/internal/auth"
	"net/http"
	"strings"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		_, claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token payload"})
			return
		}
		role, _ := claims["role"].(string)
		c.Set("userID", uint(userIDFloat))
		c.Set("role", role)
		c.Next()
	}
}
