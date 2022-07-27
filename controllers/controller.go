package controllers

import (
	"github.com/hermannhahn/go-api/database"
	"github.com/hermannhahn/go-api/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api

// ShowProducts godoc
// @Summary List all products
// @Description returns message and a list of products
// @Tags /products
// @Produce json
// @Success 200 {object} models.ResponseList
// @Router /products [get]
func ShowProducts(c *gin.Context) {
	products := models.Products{}
	database.Connect()
	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"message": "Returning all products",
		"data":    products})
}

// ShowProduct godoc
// @Summary Get product by ID
// @Description returns message and a product
// @Tags /products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products/{id} [get]
func ShowProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	database.Connect()
	database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Returning product data for id: " + id,
		"data":    product})
}

// SearchProducts godoc
// @Summary Search products by name, description, category or price
// @Description returns message and a list of products
// @Tags /products
// @Accept json
// @Produce json
// @Param query path string true "Search term"
// @Success 200 {object} models.ResponseList
// @Failure 400 {object} models.ResponseList
// @Router /products/s/{query} [get]
func SearchProducts(c *gin.Context) {
	products := models.Products{}
	search := c.Param("query")
	database.Connect()
	database.DB.Where("name LIKE ?", "%"+search+"%").Find(&products)
	if len(products) == 0 {
		database.DB.Where("description LIKE ?", "%"+search+"%").Find(&products)
		if len(products) == 0 {
			database.DB.Where("category LIKE ?", "%"+search+"%").Find(&products)
			if len(products) == 0 {
				database.DB.Where("price LIKE ?", "%"+search+"%").Find(&products)
				if len(products) == 0 {
					c.JSON(http.StatusNotFound, gin.H{"error": "Product(s) not found"})
					return
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Returning results for search term: " + search,
		"data":    products})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description creates a new product
// @Tags /products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidadeProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.Connect()
	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "New product created",
		"data":    product})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description deletes a product
// @Tags /products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	database.Connect()
	database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	database.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product ID: " + id + " deleted",
		"data":    product})
}

// UpdateProduct godoc
// @Summary Update a product
// @Description updates a product
// @Tags /products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /products/{id} [patch]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	database.Connect()
	database.DB.First(&product, id)
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidadeProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Product ID: " + id + " updated",
		"data":    product})
}

// ShowIndex is the index page
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Go API",
		"welcome": "Welcome to the Go API!",
	})
}
