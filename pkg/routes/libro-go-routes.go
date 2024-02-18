package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marktlinn/libro-go/pkg/config"
	"github.com/marktlinn/libro-go/pkg/controllers"
)

var RegisterLibroGoRoutes = func(router *mux.Router) {
	dbConnection, err := config.Connect()
	if err != nil {
		panic(err)
	}
	db := dbConnection.GetDB()

	router.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateBook(w, r, db)
	},
	).Methods("POST")

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetBooks(w, r, db)
	},
	).Methods("GET")

	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetBookById(w, r, db)
	},
	).Methods("GET")

	router.HandleFunc("/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateBook(w, r, db)
	},
	).Methods("PUT")

	router.HandleFunc("/book/{id}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteBook(w, r, db)
	},
	).Methods("DELETE")
}
