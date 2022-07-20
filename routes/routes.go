package routes

import (
	"github.com/hermannhahn/go-api-gin/controllers"
	"github.com/hermannhahn/go-api-gin/database"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// HandleRequests is a function that handles all requests
func HandleRequests() {
	database.Connect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/", controllers.ShowIndex)
	r.GET("/products", controllers.ShowProducts)
	r.GET("/products/s/:query", controllers.SearchProducts)
	r.GET("/products/:id", controllers.ShowProduct)
	r.POST("/products", controllers.CreateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.Run(":8080")
}
