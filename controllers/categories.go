package controllers

import (
	"github.com/hermannhahn/go-api/database"
	"github.com/hermannhahn/go-api/models"
)

// DON'T CHANGE FUNCTION COMMENTS IF YOU DON'T KNOW WHAT YOU ARE DOING. IT'S USED BY SWAGGER TO GENERATE DOCUMENTATION.
// SWAGGER DOCUMENTATION: https://github.com/swaggo/gin-swagger

// @BasePath /api

// GetCategoryName returns the name of the category
func GetCategoryName(id string) string {
	category := models.Category{}
	database.Connect()
	database.DB.First(&category, id)
	if category.ID == 0 {
		return "Unknown"
	}
	return category.Name
}
