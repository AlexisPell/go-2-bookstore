package routes

import (
	"net/http"

	"github.com/alexispell/go-2-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	// POST
	router.HandleFunc("/books/", controllers.CreateBook).Methods(http.MethodPost)
	// GET
	router.HandleFunc("/books/", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	// PUT
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	// DELETE
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods(http.MethodDelete)
}
