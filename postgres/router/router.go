package routes

import (
	
	"postgres/middleware"

	"github.com/gorilla/mux"
)

func RoutesControl(router *mux.Router) {

	// router.HandleFunc("/books", controller.GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE")

	router.HandleFunc("/api/newstock", middleware.GetAllStock).Methods("GET")
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET")
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST")
	
}