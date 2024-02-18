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

	r := mux.NewRouter()
	routes.RegisterLibroGoRoutes(r)
	http.Handle("/", r)

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Server running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
