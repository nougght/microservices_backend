package middleware

import (
	"net/http"
	"store-server/internal/auth/tools"
	"strings"

	"github.com/gin-gonic/gin"
)

func AccessTokenAuthMiddleware(jwtTools *tools.JwtTools) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")

		userId, err := jwtTools.ParseJWTToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userId", userId)
	}
}
