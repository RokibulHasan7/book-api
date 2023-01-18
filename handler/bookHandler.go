package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/RokibulHasan7/book-api/model"

	"github.com/RokibulHasan7/book-api/db"
)

func parseURL(url string) string {
	param := strings.Split(url, "/")
	return param[len(param)-1]
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Books)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newBook model.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "Can not decode data", http.StatusBadRequest)
		return
	}
	if len(newBook.Id) == 0 || len(newBook.Name) == 0 || len(newBook.ISBN) == 0 || len(newBook.PublishDate) == 0 || len(newBook.AuthorList) == 0 {
		http.Error(w, "Book Id or Name or ISBN or Publish date can not Null", http.StatusBadRequest)
		return
	}

	if len(newBook.AuthorList[0].FirstName) == 0 || len(newBook.AuthorList[0].LastName) == 0 || len(newBook.AuthorList[0].Email) == 0 {
		http.Error(w, "Author name or Email can not Null", http.StatusBadRequest)
		return
	}

	db.Books = append(db.Books, newBook)
	db.BookMap[newBook.ISBN] = newBook

	//json.NewEncoder(w).Encode(newBook)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book added successfully."))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var updatedBook model.Book
	idParam := parseURL(r.URL.Path)
	json.NewDecoder(r.Body).Decode(&updatedBook)
	/*updateBook := model.Book{request["id"],
	request["name"], request["isbn"],
	request["publishDate"],
	[]model.Author{{request["firstname"],
		request["lastname"], request["email"],
	}}}*/

	// Checking information of updatedBook
	if len(updatedBook.Id) == 0 || len(updatedBook.Name) == 0 || len(updatedBook.ISBN) == 0 || len(updatedBook.PublishDate) == 0 || len(updatedBook.AuthorList) == 0 {
		http.Error(w, "Book Id or Name or ISBN or Publish date can not Null", http.StatusBadRequest)
		return
	}

	if len(updatedBook.AuthorList[0].FirstName) == 0 || len(updatedBook.AuthorList[0].LastName) == 0 || len(updatedBook.AuthorList[0].Email) == 0 {
		http.Error(w, "Author name or Email can not Null", http.StatusBadRequest)
		return
	}

	check := false
	for i, bookVal := range db.Books {
		if bookVal.Id == idParam {
			db.Books[i] = updatedBook
			check = true
			break
		}
	}
	if !check {
		w.WriteHeader(404)
		w.Write([]byte("Id not Exists."))
		return
	}

	//json.NewEncoder(w).Encode(updateBook)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book updated successfully."))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	idParam := parseURL(r.URL.Path)
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
		w.WriteHeader(404)
		w.Write([]byte("Book not found."))
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(returnBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	idParam := parseURL(r.URL.Path)
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
		w.WriteHeader(404)
		w.Write([]byte("Profile Not Found."))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Profile Deleted."))
}
