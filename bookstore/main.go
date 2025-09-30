package main

import (
	"bookstore/controller"
	"bookstore/db"
	"bookstore/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func RoutesControl(router *mux.Router) {

	router.HandleFunc("/books", controller.GetAllBooksHandler).Methods("GET")
	// router.HandleFunc("/books/{id}", controller.CreateBookHandler).Methods("GET")
	router.HandleFunc("/books", controller.CreateManyBooksHandler).Methods("POST")
	// router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")

}

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
	client, _, err := db.ConnectDB(mongoURI, dbName)
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	// Initialize collections
	// controller.InitCollections(client, database)
	model.InitCollections(client, dbName)
	// Create router
	r := mux.NewRouter()
	RoutesControl(r)
	// Register routes
	// controller.RegisterRoutes(r)

	// Start server
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
