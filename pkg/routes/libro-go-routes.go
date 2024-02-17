package routes

import (
	"github.com/gorilla/mux"
	"github.com/marktlinn/libro-go/pkg/controllers"
)

var RegisterLibroGoRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

	router.HandleFunc("/books/", controllers.GetBook).Methods("GET")

	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")

	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")

	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
