// handlers/auth_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce-api/database"
	"ecommerce-api/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{Name: req.Name, Email: req.Email, Password: string(hashedPassword)}
	database.DB.Create(&user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
