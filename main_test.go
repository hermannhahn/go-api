package main

import (
	"go-api-gin/controllers"
	"go-api-gin/database"
	"go-api-gin/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

// SetupTestRoutes is a function that returns a router with all routes
func SetupTestRoutes(route string) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	database.Connect()
	if route == "products" {
		r.GET("/products", controllers.ShowProducts)
	}
	if route == "product" {
		r.GET("/products/:id", controllers.ShowProduct)
	}
	if route == "create" {
		r.POST("/products", controllers.CreateProduct)
	}
	if route == "update" {
		r.PATCH("/products/:id", controllers.UpdateProduct)
	}
	if route == "delete" {
		r.DELETE("/products/:id", controllers.DeleteProduct)
	}
	if route == "search" {
		r.GET("/products/s/:query", controllers.SearchProducts)
	}
	return r
}

func TestStatusOK(t *testing.T) {
	r := SetupTestRoutes("products")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
}

func TestCreateMockProduct(t *testing.T) {
	SetupTestRoutes("create")
	product := models.Product{Name: "TestingCreateProduct", Description: "Test Description", Price: 1.1, Quantity: 1, Image: "http://test.com/test.jpg", Active: true}
	database.DB.Create(&product)
	ID = 0
	assert.NotEqual(t, 0, product.ID, "Product not created. ID should not be 0")
	ID = int(product.ID)
}

func TestUpdateMockProduct(t *testing.T) {
	SetupTestRoutes("update")
	var product models.Product
	database.DB.First(&product, ID)
	product.Name = "TestingUpdateProduct"
	database.DB.Save(&product)
	assert.Equal(t, "TestingUpdateProduct", product.Name, "Product not updated")
}

func TestSearchMockProduct(t *testing.T) {
	r := SetupTestRoutes("search")
	var product models.Product
	database.DB.First(&product, ID)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/s/TestingUpdateProduct", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Status code should be 200")
}

func TestDeleteMockProduct(t *testing.T) {
	SetupTestRoutes("delete")
	var product models.Product
	notDeleted := product.DeletedAt
	database.DB.First(&product, ID)
	database.DB.Delete(&product)
	assert.NotEqual(t, notDeleted, product.DeletedAt, "Product not deleted")
}

func TestRemoveMockProductFromDatabase(t *testing.T) {
	SetupTestRoutes("delete")
	var product models.Product
	database.DB.Unscoped().First(&product, ID)
	notDeleted := product.ID
	database.DB.Unscoped().Delete(&product)
	assert.Equal(t, notDeleted, product.ID, "Product not removed from database")
}
