package routes

import (
	"bookstore/controller"

	"github.com/gorilla/mux"
)

func RoutesControl(router *mux.Router) {

	// router.HandleFunc("/books", controller.GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/{id}", controller.UpdateOneBookHandler).Methods("PUT")
	router.HandleFunc("/books", controller.DeleteAllBooksHandler).Methods("DELETE")

	router.HandleFunc("/books", controller.GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetOneBookHandler).Methods("GET")
	router.HandleFunc("/books", controller.CreateManyBooksHandler).Methods("POST")
	router.HandleFunc("/books", controller.CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", controller.DeleteOneBookHandler).Methods("DELETE")

}
