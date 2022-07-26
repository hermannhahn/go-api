package main

import (
	"github.com/hermannhahn/go-api/routes"

	docs "github.com/hermannhahn/go-api/docs"
)

// @title			Simple Rest API
// @version			1.1
// @description		This is a sample rest api for golang
// @termsOfService	http://swagger.io/terms/

// @contact.name	API Support
// @contact.url		https://github.com/hermannhahn/go-api
// @contact.email	hermann.h.hahn@gmail.com

// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html

// @host			localhost:8080
// @BasePath		/api

// @Schemes 		http
// @Produce 		json
// @Consume 		json

// @securityDefinitions.apikey  my-api-key
// @in                          header
// @name                        Authorization
// @description					API Key Authentication
func main() {
	println("API: http://localhost:8080/api/")
	println("Swagger documentation: http://localhost:8080/swagger/")
	println("API started.")
	docs.SwaggerInfo.BasePath = "/"
	routes.HandleRequests()
}
