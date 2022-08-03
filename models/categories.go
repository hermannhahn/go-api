package models

import (
	"gorm.io/gorm"
)

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
