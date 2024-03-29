package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Authorization is a middleware that validates the authentication
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get the authorization header sent by the client
		token := c.GetHeader("Authorization")
		// Get the key from the environment variable
		apiKey := os.Getenv("API_KEY")
		// Check if debug mode is enabled
		debug := os.Getenv("DEBUG")

		// Disable authorization if debug mode is enabled
		if debug == "false" {
			if len(token) == 0 {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization is required Header"})
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if token != apiKey {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
					c.AbortWithStatus(http.StatusUnauthorized)
				}
			}
		}

		// Hand over to the next middleware
		c.Next()
	}
}
