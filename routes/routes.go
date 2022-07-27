package routes

import (
	"github.com/hermannhahn/go-api/controllers"
	"github.com/hermannhahn/go-api/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/index.html")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api")
	{
		api.GET("", controllers.ShowIndex)
		api.Use(middleware.Authorization(), middleware.CORS())
		products := api.Group("/products")
		{
			products.GET("", controllers.ShowProducts)
			products.GET("/s/:query", controllers.SearchProducts)
			products.GET("/:id", controllers.ShowProduct)
			products.POST("", controllers.CreateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
			products.PATCH("/:id", controllers.UpdateProduct)
		}
	}

	r.Run(":8080")
}
