package routes

import (
	"os"

	"github.com/hermannhahn/go-api/controllers"
	"github.com/hermannhahn/go-api/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {

	// Get debug mode from environment variable
	debug := os.Getenv("DEBUG")

	// Set gin mode
	if debug == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new gin router
	r := gin.New()

	r.SetTrustedProxies([]string{os.Getenv("API_TRUSTED_PROXIES")})

	// Create APP index page
	r.LoadHTMLGlob("templates/index.html")

	// Load assets
	r.Static("/assets", "./assets")

	// Create API documentation page
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Create API routes
	api := r.Group("/api")
	{
		// route: /api
		api.GET("", controllers.ShowIndex)

		// Create a new middleware handler for API
		api.Use(middleware.Authorization(), middleware.CORS(), gin.Recovery(), gin.Logger())

		// route: /api/products
		products := api.Group("/products")
		{
			products.GET("", controllers.ShowProducts)            // Show all products
			products.GET("/s/:query", controllers.SearchProducts) // Search for a product
			products.GET("/:id", controllers.ShowProduct)         // Show a single product
			products.POST("", controllers.CreateProduct)          // Create a new product
			products.DELETE("/:id", controllers.DeleteProduct)    // Delete a product
			products.PATCH("/:id", controllers.UpdateProduct)     // Update a product
		}
	}

	// Start the server
	r.Run(":" + string(os.Getenv("API_PORT")))
}
