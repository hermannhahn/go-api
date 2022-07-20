package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Configuration for the database application
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

// Products struct
type Products []Product

// Response struct
type Response struct {
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

// ResponseList struct
type ResponseList struct {
	Message string   `json:"message"`
	Data    Products `json:"data"`
}

// ValidadeProduct validates the product
func ValidadeProduct(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}
	return nil
}
