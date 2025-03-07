// handlers/transaction_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce-api/database"
	"ecommerce-api/models"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
	var cartItems []models.CartItem
	database.DB.Preload("Product").Find(&cartItems)

	if len(cartItems) == 0 {
		http.Error(w, "keranjang kosong", http.StatusBadRequest)
		return
	}

	totalAmount := 0.0
	var transactionItems []models.TransactionItem

	for _, item := range cartItems {
		if item.Quantity > item.Product.Stock {
			http.Error(w, "Insufficient stock for product: "+item.Product.Name, http.StatusBadRequest)
			return
		}
		totalAmount += float64(item.Quantity) * item.Product.Price
		transactionItems = append(transactionItems, models.TransactionItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
		item.Product.Stock -= item.Quantity
		database.DB.Save(&item.Product)
	}

	transaction := models.Transaction{TotalAmount: totalAmount, Items: transactionItems}
	database.DB.Create(&transaction)
	database.DB.Delete(&cartItems)

	json.NewEncoder(w).Encode(transaction)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transaction
	database.DB.Preload("Items.Product").Find(&transactions)
	json.NewEncoder(w).Encode(transactions)
}
