package main

import (
	"github.com/gorilla/mux"
	"github.com/iamadarsh0309/Kindle-Go/pkg/routes"

	// "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
