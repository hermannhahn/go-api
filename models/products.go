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

// Products struct
type Products []Product

// ValidadeProduct validates the product
func ValidadeProduct(product *Product) error {
	if err := validator.Validate(product); err != nil {
		return err
	}
	return nil
}
