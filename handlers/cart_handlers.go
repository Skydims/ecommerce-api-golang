// handlers/cart_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce-api/database"
	"ecommerce-api/models"
	"github.com/gorilla/mux"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cartItem models.CartItem
	json.NewDecoder(r.Body).Decode(&cartItem)

	// Cek apakah produk ada
	var product models.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Tambahkan produk ke keranjang
	cartItem.Product = product
	database.DB.Create(&cartItem)
	json.NewEncoder(w).Encode(cartItem)
}

func GetCartItems(w http.ResponseWriter, r *http.Request) {
	var cartItems []models.CartItem
	database.DB.Preload("Product").Find(&cartItems)
	json.NewEncoder(w).Encode(cartItems)
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var cartItem models.CartItem
	if err := database.DB.First(&cartItem, params["id"]).Error; err != nil {
		http.Error(w, "Cart item not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&cartItem)
	w.WriteHeader(http.StatusNoContent)
}