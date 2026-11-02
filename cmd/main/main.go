package main

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/iamadarsh0309/Kindle-Go/pkg/routes"

	// "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "localhost:9010"
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(addr, r))
}
