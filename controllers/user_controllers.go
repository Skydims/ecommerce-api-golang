package controllers

import (
	"encoding/json"
	"net/http"

	"ecommerce-api/database"
	"ecommerce-api/models"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// GetProfileHandler mengambil data profil user dari token JWT
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == authHeader {
		http.Error(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
		return
	}

	// Parse token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrNoCookie
		}
		return []byte("SECRET_KEY_KAMU"), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// Ambil email dari claims JWT
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
		return
	}

	email, ok := claims["email"].(string)
	if !ok {
		http.Error(w, "Unauthorized: Email not found in token", http.StatusUnauthorized)
		return
	}

	// Cari user berdasarkan email
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kirim response user data
	response := map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
