package models

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
	ID          string
	Name        string
	Description string
	Image       string
	Price       float64
	Quantity    int
	Active      bool
}

// Products is a slice of Product
type Products []Product // slice of products
