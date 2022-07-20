// dev: github.com/hermannhahn
package main

import (
	"github.com/hermannhahn/go-api-gin/routes"

	docs "github.com/hermannhahn/go-api-gin/docs"
)

// @title Simple Rest API
// @version 1.0
// @description This is a sample product store database api.

// @contact.name Hermann Hahn
// @contact.url https://github.com/hermannhahn/go-api-gin
// @contact.email hermann.h.hahn@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	println("Starting API...")
	println("Swagger documentation: http://localhost:8080/swagger/")
	println("API documentation: http://localhost:8080/api/")
	docs.SwaggerInfo.BasePath = "/"
	routes.HandleRequests()
	println("API started.")
}
