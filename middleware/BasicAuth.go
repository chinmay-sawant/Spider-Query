package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header value
		if c.Request.URL.Path == "/login" || strings.HasPrefix(c.Request.URL.Path, "/login/") || strings.HasPrefix(c.Request.URL.Path, "/view/") {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")

		// Check if the header is present
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Parse the Basic Auth header
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Basic" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Decode the base64 encoded credentials
		decoded, err := base64.StdEncoding.DecodeString(authParts[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Split the decoded credentials into username and password
		credentials := strings.Split(string(decoded), ":")
		if len(credentials) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the username and password match
		if credentials[0] != username || credentials[1] != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If authentication is successful, proceed to the next middleware or handler
		c.Next()
	}
}
