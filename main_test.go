package main

import (
	"go-api-gin/controllers"
	"go-api-gin/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// SetupTestRoutes is a function that returns a router with all routes
func SetupTestRoutes() *gin.Engine {
	database.Connect()
	r := gin.Default()
	r.GET("/products", controllers.ShowProducts)
	r.GET("/products/s/:query", controllers.SearchProducts)
	r.GET("/products/:id", controllers.ShowProduct)
	r.POST("/products", controllers.CreateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	return r
}

func TestStatusOK(t *testing.T) {
	r := SetupTestRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}
}
