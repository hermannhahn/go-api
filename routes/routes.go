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
	r.GET("/products/all", controllers.ShowProducts)
	r.GET("/products/search/:query", controllers.SearchProducts)
	r.GET("/product/show/:id", controllers.ShowProduct)
	r.POST("/product/new", controllers.CreateProduct)
	r.DELETE("/product/delete/:id", controllers.DeleteProduct)
	r.PATCH("/product/update/:id", controllers.UpdateProduct)
	r.Run(":8080")
}
