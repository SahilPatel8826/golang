package controller

import (
	"bookstore/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// Package-level variable to store the books collection
var BookCollection *mongo.Collection

// InitCollections initializes MongoDB collections
// func InitCollections(client *mongo.Client, database *mongo.Database) {

// 	BookCollection = database.Collection("books")
// }

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := model.GetAllBooks()
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
	var books []model.Book

	// Decode JSON into []Book
	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Insert multiple books
	insertedIDs, err := model.CreateManyBooks(books)
	if err != nil {
		http.Error(w, "Failed to insert books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with inserted IDs
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertedIDs)
}

func DeleteOneBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Call repo function
	model.DeleteOneBook(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Book deleted successfully"})

}
func GetOneBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	result, err := model.GetOneBook(id)
	if err != nil {
		http.Error(w, "Book not found: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
	// fmt.Printf("Fetched book: %+v\n", result)
}

func DeleteAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deletedCount, err := model.DeleteAllBooks()
	if err != nil {
		http.Error(w, "Failed to delete books: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":       "All books deleted successfully",
		"deleted_count": deletedCount,
	})
}
func UpdateOneBookHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call model update function
	updatedBook, err := model.UpdateOne(id, book)
	if err != nil {
		http.Error(w, "Failed to update book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Book updated successfully",
		"book":    updatedBook,
	})
}
