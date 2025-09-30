package main

import (
	"bookstore/controller"

	"github.com/gorilla/mux"
)

func RoutesControl(router *mux.Router) {

	router.HandleFunc("/books", controller.GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/books", controller.CreateManyBooksHandler).Methods("POST")
	router.HandleFunc("/book", controller.CreateBookHandler).Methods("POST")
	// router.HandleFunc("/books/{id}", controller.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", controller.DeleteBook).Methods("DELETE")

}
