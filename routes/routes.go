package routes

import (
	"go-api-gin/controllers"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	router := gin.Default()
	router.GET("/products/all", controllers.ShowProducts)
	router.GET("/products/show/:id", controllers.ShowProduct)
	router.GET("/products/search/:query", controllers.SearchProducts)
	router.Run(":8080")
}
