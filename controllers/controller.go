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
