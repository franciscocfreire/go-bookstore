package routes

import (
	"github.com/franciscocfreire/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/boook/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
