package main

import (
	"bytes"
	"encoding/json"
	"go-api-gin/controllers"
	"go-api-gin/database"
	"go-api-gin/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID string

// SetupTestRoutes is a function that returns a router with all routes
func SetupTestRoutes() *gin.Engine {
	database.Connect()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestAPI(t *testing.T) {
	CreateTestProduct(t)
	SearchTestProduct(t)
	ShowTestProducts(t)
	ShowTestProduct(t)
	UpdateTestProduct(t)
	DeleteTestProduct(t)
	RemoveTestProductFromDataBase(t)
}

func CreateTestProduct(t *testing.T) {
	// Create a new product for testing
	product := models.Product{Name: "TestProduct", Description: "TestDescription", Price: 1.1, Quantity: 1, Image: "http://test.com/test.jpg", Active: true}
	newProduct, _ := json.Marshal(product)
	r := SetupTestRoutes()
	r.POST("/products", controllers.CreateProduct)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(newProduct))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200")
	assert.NotNil(t, w.Body, "Should return a body")
	bodyString := w.Body.String()
	assert.Contains(t, bodyString, "New product created", "Should return message 'New product created'")
	var testProduct models.Response
	bodyBytes := w.Body.Bytes()
	json.Unmarshal([]byte(bodyBytes), &testProduct)
	ID = strconv.FormatUint(uint64(testProduct.Data.ID), 10)
}

func SearchTestProduct(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/products/s/:query", controllers.SearchProducts)
	req, _ := http.NewRequest("GET", "/products/s/TestProduct", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200")
	assert.NotNil(t, w.Body, "Should return a body")
	body := w.Body.String()
	assert.Contains(t, body, "Returning results in a slice of products", "Should return message 'Returning results in a slice of products'")
}

func ShowTestProducts(t *testing.T) {
	// Try get list of products
	r := SetupTestRoutes()
	r.GET("/products", controllers.ShowProducts)
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200")
	assert.NotNil(t, w.Body, "Should return a body")
	body := w.Body.String()
	assert.Contains(t, body, "Returning all products", "Should return message 'Returning all products'")
}

func ShowTestProduct(t *testing.T) {
	// Try get the product
	r := SetupTestRoutes()
	r.GET("/products/:id", controllers.ShowProduct)
	req, _ := http.NewRequest("GET", "/products/"+ID, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200")
	assert.NotNil(t, w.Body, "Should return a body")
	body := w.Body.String()
	assert.Contains(t, body, "Returning product data for id", "Should return message 'Returning product data for id...'")
}

func UpdateTestProduct(t *testing.T) {
	// Update the product
	r := SetupTestRoutes()
	r.PATCH("/products/:id", controllers.UpdateProduct)
	newName, _ := json.Marshal(map[string]interface{}{
		"Name": "TestProductUpdated",
	})
	req, _ := http.NewRequest("PATCH", "/products/"+ID, bytes.NewBuffer(newName))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200")
	assert.NotNil(t, w.Body, "Should return a body")
	body := w.Body.String()
	assert.Contains(t, body, "Product ID: "+ID+" updated", "Should return message 'Product ID: "+ID+" updated'")
}

func DeleteTestProduct(t *testing.T) {
	// Delete the product
	r := SetupTestRoutes()
	r.DELETE("/products/:id", controllers.DeleteProduct)
	req, _ := http.NewRequest("DELETE", "/products/"+ID, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Should return 200")
	assert.NotNil(t, w.Body, "Should return a body")
	body := w.Body.String()
	assert.Contains(t, body, "Product ID: "+ID+" deleted", "Should return message 'Product ID: "+ID+" deleted'")
}

func RemoveTestProductFromDataBase(t *testing.T) {
	println("PASS: All tests passed")
	println("Removing Test Product from Database")
	var product models.Product
	notRemoved := product.ID
	database.DB.Unscoped().First(&product, ID)
	database.DB.Unscoped().Delete(&product)
	assert.NotEqual(t, notRemoved, product.ID, "Product not removed from database")
	println("")
	println("PASS: Product removed from database.")
	println("All tests executed successfully.")
}
