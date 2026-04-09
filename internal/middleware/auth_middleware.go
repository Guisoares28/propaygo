package middleware

import (
	"net/http"
	"propay/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenService *service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error: ": "Token não informado",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := tokenService.DecodedToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token inválido ou Expirado",
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
