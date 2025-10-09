package routes

import (
	"postgres/handler"

	"github.com/gorilla/mux"
)

func RoutesControl(router *mux.Router) {

	// // router.HandleFunc("/books", controller.GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/api/stock/{id}", handler.UpdateStockHandler).Methods("PUT")
	router.HandleFunc("/api/stock/{id}", handler.DeleteStockHandler).Methods("DELETE")

	router.HandleFunc("/api/newstock", handler.GetAllStockHandler).Methods("GET")
	router.HandleFunc("/api/stock/{id}", handler.GetStockHandler).Methods("GET")
	router.HandleFunc("/api/newstock", handler.CreateStockHandler).Methods("POST")

}
