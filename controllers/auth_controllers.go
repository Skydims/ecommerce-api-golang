package controllers

import (
	"encoding/json"
	"net/http"


	"ecommerce-api/database"
	"ecommerce-api/middleware"
	"ecommerce-api/models"

	"gorm.io/gorm"
)

// Struct untuk request login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler menangani proses login dan mengembalikan token JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Cari user di database berdasarkan email
	var user models.User
	result := database.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "User not found", http.StatusUnauthorized)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Cek apakah password cocok (sementara tanpa hashing, harusnya pakai bcrypt)
	if user.Password != req.Password {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Buat token JWT
	token, err := middlewares.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Kirim response token
	response := map[string]string{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
