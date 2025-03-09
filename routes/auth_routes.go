// routes/auth_routes.go
package routes

import (
	"ecommerce-api/handlers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", handlers.Register).Methods("POST")
}