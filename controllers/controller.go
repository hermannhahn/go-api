package controllers

import (
	"go-api-gin/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowProducts is a function that returns a list of products
func ShowProducts(c *gin.Context) {
	db := data.Connect()
	c.JSON(http.StatusOK, data.ShowProducts(db))
}

// ShowProduct is a function that returns a product
func ShowProduct(c *gin.Context) {
	db := data.Connect()
	c.JSON(http.StatusOK, data.ShowProduct(db, c.Param("id")))
}

// SearchProducts is a function that returns a list of products with a search term
func SearchProducts(c *gin.Context) {
	db := data.Connect()
	c.JSON(http.StatusOK, data.SearchProducts(db, c.Param("query")))
}
