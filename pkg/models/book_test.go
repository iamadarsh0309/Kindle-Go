package models

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var reqCounter atomic.Int64

// uniqueISBN keeps repeated test runs independent of existing rows.
func uniqueISBN() string {
	return fmt.Sprintf("TEST-%d", reqCounter.Add(1))
}

func TestCreateBook(t *testing.T) {
	book, err := CreateBook(Bookify{Isbn: uniqueISBN(), Title: "Go Test"})
	if err != nil {
		t.Fatalf("CreateBook failed: %v", err)
	}
	if book.ID == 0 {
		t.Fatalf("expected generated ID, got %d", book.ID)
	}
}

func TestGetAllBooks(t *testing.T) {
	b1, err := CreateBook(Bookify{Isbn: uniqueISBN(), Title: "All-1"})
	if err != nil {
		t.Fatalf("CreateBook failed: %v", err)
	}
	books, err := GetAllBooks()
	if err != nil {
		t.Fatalf("GetAllBooks failed: %v", err)
	}
	found := false
	for _, b := range books {
		if b.ID == b1.ID {
			found = true
		}
	}
	if !found {
		t.Errorf("created book %+v missing from GetAllBooks result (%d books)", b1, len(books))
	}
}

func TestGetBookById(t *testing.T) {
	created, err := CreateBook(Bookify{Isbn: uniqueISBN(), Title: "Find Me"})
	if err != nil {
		t.Fatalf("CreateBook failed: %v", err)
	}

	got, err := GetBookById(fmt.Sprint(created.ID))
	if err != nil {
		t.Fatalf("GetBookById failed: %v", err)
	}
	if got.ID != created.ID || got.Title != created.Title {
		t.Errorf("got %+v, want %+v", got, created)
	}
}

func TestGetBookByIdNotFound(t *testing.T) {
	if _, err := GetBookById("999999999"); err == nil {
		t.Error("expected error for non-existent ID, got nil")
	}
}

func TestUpdateBookById(t *testing.T) {
	created, err := CreateBook(Bookify{Isbn: uniqueISBN(), Title: "Before"})
	if err != nil {
		t.Fatalf("CreateBook failed: %v", err)
	}

	updated, err := UpdateBookById(fmt.Sprint(created.ID), Bookify{Isbn: created.Isbn, Title: "After"})
	if err != nil {
		t.Fatalf("UpdateBookById failed: %v", err)
	}
	if updated.Title != "After" {
		t.Errorf("expected title After, got %+v", updated)
	}
}

func TestDeleteBookById(t *testing.T) {
	created, err := CreateBook(Bookify{Isbn: uniqueISBN(), Title: "Doomed"})
	if err != nil {
		t.Fatalf("CreateBook failed: %v", err)
	}

	deleted, err := DeleteBookById(fmt.Sprint(created.ID))
	if err != nil {
		t.Fatalf("DeleteBookById failed: %v", err)
	}
	if deleted.ID != created.ID {
		t.Errorf("deleted wrong book: got ID %d, want %d", deleted.ID, created.ID)
	}
	if _, err := GetBookById(fmt.Sprint(created.ID)); err == nil {
		t.Error("book still fetchable after delete")
	}
}
