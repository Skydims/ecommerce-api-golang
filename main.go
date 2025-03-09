// main.go
package main

import (
	"log"
	"net/http"

	"ecommerce-api/database"
	"ecommerce-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB() // Inisialisasi database

	r := mux.NewRouter()
	routes.RegisterProductRoutes(r)
	routes.RegisterCartRoutes(r)
	routes.RegisterTransactionRoutes(r)
        routes.RegisterAuthRoutes(r)
	
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
