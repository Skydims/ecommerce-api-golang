// models/product.go
package models

import "gorm.io/gorm"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	gorm.Model
}
