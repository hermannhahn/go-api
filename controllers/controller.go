package controllers

import (
	"go-api-gin/database"
	"go-api-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowProducts is a function that returns a list of products
func ShowProducts(c *gin.Context) {
	products := models.Products{}
	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"message": "Returning all products",
		"data":    products})
}

// ShowProduct is a function that returns a product
func ShowProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	database.DB.First(&product, id)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Returning product data for id: " + id,
		"data":    product})
}

// SearchProducts is a function that returns a list of products with a search term [name or description or price]
func SearchProducts(c *gin.Context) {
	products := models.Products{}
	search := c.Param("query")
	database.DB.Where("name LIKE ?", "%"+search+"%").Find(&products)
	if len(products) == 0 {
		database.DB.Where("description LIKE ?", "%"+search+"%").Find(&products)
		if len(products) == 0 {
			database.DB.Where("price LIKE ?", "%"+search+"%").Find(&products)
			if len(products) == 0 {
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Returning results in a slice of products",
		"data":    products})
}

// CreateProduct is a function that creates a product
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
	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "New product created",
		"data":    product})
}

// DeleteProduct is a function that deletes a product
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
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

// UpdateProduct is a function that updates a product
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
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
