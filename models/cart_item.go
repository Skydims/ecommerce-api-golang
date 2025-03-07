// models/cart_item.go
package models

import "gorm.io/gorm"

type CartItem struct {
	ID        uint    `gorm:"primaryKey"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int     `json:"quantity"`
	gorm.Model
}
