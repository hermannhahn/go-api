package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	CategoryID  uint    `json:"category_id"`
	Name        string  `json:"name" validate:"nonzero"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"nonzero"`
	Quantity    int     `json:"quantity"`
	Active      bool    `json:"active"`
}

// Category struct
type Category struct {
	gorm.Model
	Parent      uint   `json:"parent"`
	Name        string `json:"name" validate:"nonzero"`
	Description string `json:"description"`
	Image       string `json:"image" validate:"regexp=^(http|https):\\/\\/.*$"`
}

// Categories struct
type Categories []Category

// Image struct
type Image struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Image     string `json:"image" validate:"regexp=^(http|https):\\/\\/.*$"`
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
