package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Configuration is the configuration for the database application
type Configuration struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
}

// Product struct
type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"nonzero"`
	Description string  `json:"description"`
	Image       string  `json:"image" validate:"regexp=^(http|https):\\/\\/.*$"`
	Price       float64 `json:"price" validate:"nonzero"`
	Quantity    int     `json:"quantity"`
	Active      bool    `json:"active"`
}

// Products is a slice of Product
type Products []Product // slice of products

// Response is the response for the API
type Response struct {
	message string
	data    []Product
}

// ValidadeProduct validates the product
func ValidadeProduct(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}
	return nil
}
