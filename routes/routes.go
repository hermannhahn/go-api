package routes

import (
	"go-api-gin/controllers"
	"go-api-gin/database"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	database.Connect()
	router := gin.Default()
	router.GET("/products/all", controllers.ShowProducts)
	router.GET("/products/search/:query", controllers.SearchProducts)
	router.GET("/product/show/:id", controllers.ShowProduct)
	router.POST("/product/new", controllers.CreateProduct)
	router.DELETE("/product/delete/:id", controllers.DeleteProduct)
	router.POST("/product/update/:id", controllers.UpdateProduct)
	router.Run(":8080")
}
