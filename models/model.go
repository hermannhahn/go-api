package models

import "gorm.io/gorm"

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
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Active      bool    `json:"active"`
}

// Products is a slice of Product
type Products []Product // slice of products
