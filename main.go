package main

import (
	docs "github.com/hermannhahn/go-api/docs"
	"github.com/hermannhahn/go-api/routes"
)

var version string
var buildDate string

// @BasePath /api
// @host localhost:8080

// @title Simple Rest API
// @version 1.1
// @description This is a simple rest api with golang

// @contact.name API Support
// @contact.url https://github.com/hermannhahn/go-api
// @contact.email hermann.h.hahn@gmail.com
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	println("")
	println("   ██████╗  ██████╗      █████╗ ██████╗ ██╗ ")
	println("  ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║ ")
	println("  ██║  ███╗██║   ██║    ███████║██████╔╝██║ ")
	println("  ██║   ██║██║   ██║    ██╔══██║██╔═══╝ ██║ ")
	println("  ╚██████╔╝╚██████╔╝    ██║  ██║██║     ██║ ")
	print("   ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝ ")
	println(version)
	println("Build Date: ", buildDate)
	println("")
	println("API: http://localhost:8080/api/")
	println("Swagger documentation: http://localhost:8080/swagger/")
	println("")
	docs.SwaggerInfo.BasePath = "/api"
	routes.HandleRequests()
}
