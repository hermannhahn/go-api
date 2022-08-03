package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// ShowIndex is the index page
func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": os.Getenv("APP_NAME"),
	})
}
