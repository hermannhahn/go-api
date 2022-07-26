package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hermannhahn/go-api/database"
)

// Authorization is a middleware that validates the authentication
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != database.APIConfig().APIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid API key"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Hand over to the next middleware
		c.Next()
	}
}

// CORS is a middleware that allows CORS
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add CORS headers
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "PATCH, POST, GET, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")
		// Stop here if its a OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		if c.Request.Header.Get("Origin") != database.APIConfig().Address+":"+database.APIConfig().Port {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Hand over to the next middleware
		c.Next()
	}
}
