package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS is a middleware that allows CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Add CORS headers
		c.Header("Access-Control-Allow-Methods", "PATCH, POST, GET, DELETE")
		c.Header("Authorization", "apiKey")

		// Stop here if its a OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// Hand over to the next middleware
		c.Next()
	}
}
