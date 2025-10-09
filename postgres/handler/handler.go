package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgres/model"

	"github.com/gorilla/mux"
)

func CreateStockHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Decode JSON request into Stock struct
	var stock model.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Step 2: Call model function to insert stock record
	err := model.CreateStock(stock)
	if err != nil {
		http.Error(w, "Failed to create stock: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Step 3: Respond to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "Stock created successfully"}`)
}
func GetAllStockHandler(w http.ResponseWriter, r *http.Request) {

	result, err := model.GetAllStocks()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
func GetStockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	result, err := model.GetStocks(id)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeleteStockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	model.DeleteStock(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("deleted")
}
func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the ID from the URL
	params := mux.Vars(r)
	id := params["id"]

	// Decode request body
	var updatedStock model.Stock
	if err := json.NewDecoder(r.Body).Decode(&updatedStock); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the update function
	stock, err := model.UpdateStock(id, updatedStock)
	if err != nil {
		http.Error(w, "Error updating stock: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(stock)
}
