package middleware

import (
	"context"
	"log"
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
		// type Token struct {
		// 	AuthTime int64                  `json:"auth_time"`
		// 	Issuer   string                 `json:"iss"`
		// 	Audience string                 `json:"aud"`
		// 	Expires  int64                  `json:"exp"`
		// 	IssuedAt int64                  `json:"iat"`
		// 	Subject  string                 `json:"sub,omitempty"`
		// 	UID      string                 `json:"uid,omitempty"`
		// 	Firebase FirebaseInfo           `json:"firebase"`
		// 	Claims   map[string]interface{} `json:"-"`
		// }
		token, err := firebase.GetAuthClient().VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		log.Printf("Verified ID token: %v\n", token)

		c.Set("externalUserId", token.UID)
		c.Next()
	}
}
