package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamadarsh0309/Kindle-Go/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

}
