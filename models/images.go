package models

import (
	"gorm.io/gorm"
)

// Image struct
type Image struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Image     string `json:"image" validate:"regexp=^(http|https):\\/\\/.*$"`
}
