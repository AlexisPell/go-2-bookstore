package main

import (
	"log"
	"net/http"

	"github.com/alexispell/go-2-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookstoreRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
