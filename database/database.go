// database/database.go
package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"ecommerce-api/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal connect ke database")
	}

	// Auto migrate model Produk
	DB.AutoMigrate(&models.Product{}, &models.CartItem{}, &models.Transaction{}, &models.TransactionItem{})


}
