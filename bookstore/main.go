package main

import (
	"bookstore/controller"
	"bookstore/db"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	port := os.Getenv("PORT")

	// Connect to MongoDB
	client, database, err := db.ConnectDB(mongoURI, dbName)
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	// Initialize collections
	controller.InitCollections(client, database)

	// Create router
	r := mux.NewRouter()

	// Register routes
	// controller.RegisterRoutes(r)

	// Start server
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
