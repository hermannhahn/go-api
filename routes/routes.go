package routes

import (
	"github.com/hermannhahn/go-api-gin/controllers"
	"github.com/hermannhahn/go-api-gin/database"
	"github.com/hermannhahn/go-api-gin/middleware"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	database.Connect()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/index.html")
	r.Use(middleware.Authorization(), middleware.CORS())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/api", controllers.ShowIndex)
	r.GET("/api/products", controllers.ShowProducts)
	r.GET("/api/products/s/:query", controllers.SearchProducts)
	r.GET("/api/products/:id", controllers.ShowProduct)
	r.POST("/api/products", controllers.CreateProduct)
	r.DELETE("/api/products/:id", controllers.DeleteProduct)
	r.PATCH("/api/products/:id", controllers.UpdateProduct)
	r.Run(":8080")
}
