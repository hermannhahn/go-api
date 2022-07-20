package routes

import (
	"net/http"

	"github.com/hermannhahn/go-api-gin/controllers"
	"github.com/hermannhahn/go-api-gin/database"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	database.Connect()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/api", controllers.ShowIndex)
	r.GET("/api/products", controllers.ShowProducts)
	r.GET("/api/products/s/:query", controllers.SearchProducts)
	r.GET("/api/products/:id", controllers.ShowProduct)
	r.POST("/api/products", validateAPIKey(), controllers.CreateProduct)
	r.DELETE("/api/products/:id", validateAPIKey(), controllers.DeleteProduct)
	r.PATCH("/api/products/:id", validateAPIKey(), controllers.UpdateProduct)
	r.Run(":8080")
}

func validateAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := database.GetConfig("api").APIKey
		if apiKey != c.Request.Header.Get("X-API-KEY") {
			c.Header("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
			c.JSON(http.StatusOK, gin.H{
				"message": "Authorization Required",
				"data":    nil})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
