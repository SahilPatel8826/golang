package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgres/model"
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

	result, err := model.GetStocks()
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
