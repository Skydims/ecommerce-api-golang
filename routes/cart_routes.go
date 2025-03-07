// routes/cart_routes.go
package routes

import (
	"ecommerce-api/handlers"
	"github.com/gorilla/mux"
)

func RegisterCartRoutes(r *mux.Router) {
	r.HandleFunc("/cart", handlers.AddToCart).Methods("POST")
	r.HandleFunc("/cart", handlers.GetCartItems).Methods("GET")
	r.HandleFunc("/cart/{id}", handlers.RemoveFromCart).Methods("DELETE")
}