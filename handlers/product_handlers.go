// handlers/product_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce-api/database"
	"ecommerce-api/models"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	if err := database.DB.First(&product, params["id"]).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	database.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	if err := database.DB.First(&product, params["id"]).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	database.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}


func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	if err := database.DB.First(&product, params["id"]).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	database.DB.Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}
