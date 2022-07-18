package routes

import (
	"go-api-gin/controllers"
	"go-api-gin/database"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	database.Connect()
	r := gin.Default()
	r.GET("/products", controllers.ShowProducts)
	r.GET("/products/s/:query", controllers.SearchProducts)
	r.GET("/products/:id", controllers.ShowProduct)
	r.POST("/products", controllers.CreateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.Run(":8080")
}
