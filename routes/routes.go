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
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.LoadHTMLGlob("templates/index.html")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api")
	{
		api.GET("", controllers.ShowIndex)
		products := api.Group("/products")
		products.Use(middleware.Authorization(), middleware.CORS(), gin.Recovery(), gin.Logger())
		{
			products.GET("", controllers.ShowProducts)
			products.GET("/s/:query", controllers.SearchProducts)
			products.GET("/:id", controllers.ShowProduct)
			products.POST("", controllers.CreateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
			products.PATCH("/:id", controllers.UpdateProduct)
		}
	}

	r.Run(":" + string(os.Getenv("API_PORT")))
}
