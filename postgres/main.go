package main

import (
	"fmt"
	"log"
	"net/http"

	"postgres/middleware" // <-- replace 'yourproject' with your actual module name
	routes "postgres/router"

	"github.com/gorilla/mux"
)

func main() {
	// Create database connection
	db := middleware.CreateConnection()
	defer db.Close()

	r := mux.NewRouter()
	routes.RoutesControl(r)

	fmt.Println("✅ Connected to PostgreSQL successfully!")
	fmt.Println("Starting server on the port 8080")

	log.Fatal(http.ListenAndServe(":8000", r))
}
