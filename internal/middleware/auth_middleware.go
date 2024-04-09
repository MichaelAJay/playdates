package middleware

import (
	"context"
	"net/http"
	"playdates/internal/firebase"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authentication token from the request
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := firebase.GetAuthClient().VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("externalUserId", token.UID)
		c.Next()
	}
}

func isValidToken(token string) bool {
	return token == "bool"
}
