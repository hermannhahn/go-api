package main

import (
	"github.com/hermannhahn/go-api/routes"

	docs "github.com/hermannhahn/go-api/docs"
)

// @title Simple Rest API
// @version 1.0
// @description This is a sample product store database api.

// @contact.name Hermann Hahn
// @contact.url https://github.com/hermannhahn/go-api
// @contact.email hermann.h.hahn@gmail.com

// @host localhost:8080
// @BasePath /api
func main() {
	println("Starting API...")
	println("Swagger documentation: http://localhost:8080/swagger/")
	println("API: http://localhost:8080/api/")
	println("API started.")
	docs.SwaggerInfo.BasePath = "/"
	routes.HandleRequests()
}
