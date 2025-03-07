// models/transaction.go
package models

import "gorm.io/gorm"

type Transaction struct {
	ID          uint             `gorm:"primaryKey"`
	TotalAmount float64          `json:"total_amount"`
	Items       []TransactionItem `gorm:"foreignKey:TransactionID" json:"items"`
	gorm.Model
}

type TransactionItem struct {
	ID            uint    `gorm:"primaryKey"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Product       Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	gorm.Model
}
