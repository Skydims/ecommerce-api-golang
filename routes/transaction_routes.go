// routes/transaction_routes.go
package routes

import (
	"ecommerce-api/handlers"
	"github.com/gorilla/mux"
)

func RegisterTransactionRoutes(r *mux.Router) {
	r.HandleFunc("/checkout", handlers.Checkout).Methods("POST")
	r.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")
}