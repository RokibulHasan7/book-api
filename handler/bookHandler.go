package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/RokibulHasan7/book-api/model"

	"github.com/RokibulHasan7/book-api/db"
)

// Get all books
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(db.Books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Post book
func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBook model.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		http.Error(w, "Can not decode data", http.StatusBadRequest)
		return
	}

	if len(newBook.Id) == 0 || len(newBook.Name) == 0 || len(newBook.ISBN) == 0 || len(newBook.PublishDate) == 0 || len(newBook.AuthorList) == 0 {
		fmt.Println("Failed to decode data1.")
		http.Error(w, "Missing required fields.", http.StatusBadRequest)
		return
	}

	if len(newBook.AuthorList[0].FirstName) == 0 || len(newBook.AuthorList[0].LastName) == 0 || len(newBook.AuthorList[0].Email) == 0 {
		http.Error(w, "Author name or Email can not Null", http.StatusBadRequest)
		return
	}

	_, ok := db.BookMap[newBook.ISBN]
	if ok {
		http.Error(w, "ISBN can't be the same.", http.StatusBadRequest)
		return
	}

	db.Books = append(db.Books, newBook)
	db.BookMap[newBook.ISBN] = newBook

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Book added successfully."))
	err = json.NewEncoder(w).Encode(newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedBook model.Book

	idParam := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Can not decode data", http.StatusBadRequest)
		return
	}

	// Checking information of updatedBook
	check := false
	for i, bookVal := range db.Books {
		if bookVal.Id == idParam {
			if bookVal.ISBN != updatedBook.ISBN && updatedBook.ISBN != "" {
				http.Error(w, "ISBN can't change.", http.StatusBadRequest)
				return
			}
			db.Books[i] = updatedBook
			check = true
			break
		}
	}
	if !check {
		//fmt.Println("Failed at -4.")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Id not Exists."))
		return
	}

	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully."))
}

// Get book by Id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	var returnBook model.Book
	check := false
	for i, bookVal := range db.Books {
		if bookVal.Id == idParam {
			returnBook = db.Books[i]
			check = true
			break
		}
	}

	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book not found."))
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(returnBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get book by ISBN
func GetBookByIsbn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	isbnParam := chi.URLParam(r, "isbn")

	book, ok := db.BookMap[isbnParam]
	if !ok {
		http.Error(w, "Book not Found!", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Delete existing book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")
	check := false

	for i, bookVal := range db.Books {
		if bookVal.Id == idParam {
			delete(db.BookMap, bookVal.ISBN)
			db.Books[i] = db.Books[len(db.Books)-1]
			db.Books = db.Books[:len(db.Books)-1]
			check = true
			break
		}
	}
	if !check {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book Not Found."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book Deleted."))
}
