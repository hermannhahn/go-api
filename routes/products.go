package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hermannhahn/go-api/controllers"
)

// LoadProductsRoutes is a function that loads products routes
func LoadProductsRoutes(api *gin.RouterGroup) {
	// route: /api/products
	products := api.Group("/products")
	{
		products.GET("", controllers.GetProducts)                    // Get all products
		products.POST("", controllers.CreateProduct)                 // Create a new product
		products.GET("/:id", controllers.GetProduct)                 // Get a product by ID
		products.DELETE("/:id", controllers.DeleteProduct)           // Delete a product by ID
		products.PATCH("/:id", controllers.UpdateProduct)            // Update a product by ID
		products.GET("/s/:query", controllers.SearchProducts)        // Search for a product by name, description or price
		products.GET("/c/:id", controllers.SearchProductsByCategory) // Search for a product by category
	}
}
