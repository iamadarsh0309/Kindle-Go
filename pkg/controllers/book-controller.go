package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamadarsh0309/Kindle-Go/pkg/models"
	"github.com/iamadarsh0309/Kindle-Go/pkg/utils"
)

// CreateBook handles POST /book/.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var payload models.Bookify
	utils.ParseBody(r, &payload)

	book, err := models.CreateBook(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetBook handles GET /book/ and returns all books.
func GetBook(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBookById handles GET /book/{bookId}.
func GetBookById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bookId"]

	book, err := models.GetBookById(id)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles PUT /book/{bookId}.
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bookId"]

	var payload models.Bookify
	utils.ParseBody(r, &payload)

	book, err := models.UpdateBookById(id, payload)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// DeleteBook handles DELETE /book/{bookId}.
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["bookId"]

	book, err := models.DeleteBookById(id)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
