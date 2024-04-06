package main

import (
	"log"
	"net/http"

	"github.com/Darkness467/Book_Management_Store/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	// handle routing

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
