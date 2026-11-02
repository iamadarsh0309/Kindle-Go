package models

import (
	"github.com/iamadarsh0309/Kindle-Go/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Isbn  string `gorm:"size:20;not null" json:"isbn"`
	Title string `gorm:"size:100;not null" json:"title"`
}

// Bookify is the JSON payload for create/update requests.
type Bookify struct {
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
}

// Bookify converts a Book into its payload form.
func (b Book) Bookify() Bookify {
	return Bookify{Isbn: b.Isbn, Title: b.Title}
}

func init() {
	config.Connect()
	db = config.DB
	db.AutoMigrate(&Book{})
}

// CreateBook inserts a new book and returns it with its generated ID.
func CreateBook(payload Bookify) (Book, error) {
	book := Book{Isbn: payload.Isbn, Title: payload.Title}
	if err := db.Create(&book).Error; err != nil {
		return Book{}, err
	}
	return book, nil
}

// GetAllBooks returns every book in the database.
func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookById returns the book with the given ID, or nil if not found.
func GetBookById(id string) (*Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// DeleteBookById removes the book with the given ID and returns its data.
func DeleteBookById(id string) (Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return Book{}, err
	}
	if err := db.Delete(&Book{}, id).Error; err != nil {
		return Book{}, err
	}
	return book, nil
}
