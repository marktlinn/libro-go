package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/marktlinn/libro-go/pkg/routes"
)

func main() {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	url := fmt.Sprintf("%s:%s", host, port)

	r := mux.NewRouter()
	routes.RegisterLibroGoRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(url, r))
}
