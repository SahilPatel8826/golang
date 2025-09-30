package controller

import (
	"bookstore/model"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// Package-level variable to store the books collection
var BookCollection *mongo.Collection

// InitCollections initializes MongoDB collections
// func InitCollections(client *mongo.Client, database *mongo.Database) {

// 	BookCollection = database.Collection("books")
// }

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := model.GetBookCollection()
	if err != nil {
		http.Error(w, "Failed to Fetch books", http.StatusInternalServerError)
		log.Println("error fetching books:", err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	// Parse JSON body into a Book struct
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call your model function to insert the book
	createdBook, err := model.CreateBook(book)
	if err != nil {
		log.Println("Failed to create book:", err)
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	// Return created book as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBook)
}
func CreateManyBooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []interface{}

	// Decode JSON directly into []interface{}
	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Insert directly (no conversion needed)
	InsertedIDs, err := model.CreateManyBooks(books)
	if err != nil {
		http.Error(w, "Failed to insert books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with inserted IDs
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(InsertedIDs)
}
