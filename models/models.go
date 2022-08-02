package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Name        string         `json:"name" validate:"nonzero"`
	Description string         `json:"description"`
	Category    string         `json:"category" validate:"nonzero"`
	Images      []ProductImage `json:"images"`
	Price       float64        `json:"price" validate:"nonzero"`
	Quantity    int            `json:"quantity"`
	Active      bool           `json:"active"`
}

// Category struct
type Category struct {
	gorm.Model
	Name           string          `json:"name" validate:"nonzero"`
	Description    string          `json:"description"`
	CategoryImages []CategoryImage `json:"category_images"`
}

// ProductImage struct
type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Image     string `json:"image" validate:"regexp=^(http|https):\\/\\/.*$"`
}

// CategoryImage struct
type CategoryImage struct {
	gorm.Model
	CategoryID uint   `json:"category_id"`
	Image      string `json:"image" validate:"regexp=^(http|https):\\/\\/.*$"`
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
