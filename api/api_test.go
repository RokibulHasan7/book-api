package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RokibulHasan7/book-api/auth"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Method       string
	Path         string
	Body         io.Reader
	ExpectedCode int
}

type TestWithID struct {
	Method       string
	Path         string
	Body         io.Reader
	ExpectedCode int
	ID           string
}

type TestWithIsbn struct {
	Method       string
	Path         string
	Body         io.Reader
	ExpectedCode int
	Isbn         string
}

func TestGetAllBooks(t *testing.T) {
	Init()          // Initialize DB
	HandleRequest() // Expose APIs

	tests := []Test{
		{
			"GET",
			"/api/v1/books",
			nil,
			200,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		fmt.Println(res)
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedCode)
	}
}

func TestPostBook(t *testing.T) {
	Init()
	HandleRequest()

	tests := []Test{
		{
			Method: "POST",
			Path:   "/api/v1/books",
			Body: bytes.NewReader([]byte(`{
			"id":          "10",
			"name":        "Celestial Bodies",
			"isbn":        "20-20-20",
			"publishDate": "01/01/2019",
			"author": [
				{
					"firstName": "Jokha",
					"lastName":  "Alharthi",
					"email":     "jokha@gmail.com"
				}
			]
		}`)),
			ExpectedCode: 201,
		},
		{
			Method: "POST",
			Path:   "/api/v1/books",
			Body: bytes.NewReader([]byte(`{
			"id":          "1",
			"name":        "Celestial Bodies",
			"isbn":        "12-12-12",
			"publishDate": "01/01/2019",
			"author": [
				{
					"firstName": "Jokha",
					"lastName":  "Alharthi",
					"email":     "jokha@gmail.com"
				}
			]
		}`)),
			ExpectedCode: 400,
		},
		{
			Method:       "POST",
			Path:         "/api/v1/books",
			Body:         bytes.NewReader([]byte(`{"id": "10"}`)),
			ExpectedCode: 400,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Path, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")

		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("token from test:", token)
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedCode)
	}
}

func TestUpdateBook(t *testing.T) {
	Init()
	HandleRequest()

	tests := []TestWithID{
		{
			Method:       "PUT",
			Path:         "/api/v1/books/%s",
			Body:         bytes.NewReader([]byte(`{"name": "test","publishDate": "12/03/2023"}`)),
			ExpectedCode: 200,
			ID:           "1",
		},
		{
			Method: "PUT",
			Path:   "/api/v1/books/%s",
			Body: bytes.NewReader([]byte(`{
			"id":          "1",
			"name":        "Celestial Bodies",
			"isbn":        "12-12-12",
			"publishDate": "01/01/2019",
			"author": [
				{
					"firstName": "Jokha",
					"lastName":  "Alharthi",
					"email":     "jokha@gmail.com"
				}
			]
		}`)),
			ExpectedCode: 400,
			ID:           "3",
		},
		{
			Method:       "PUT",
			Path:         "/api/v1/books/%s",
			Body:         bytes.NewReader([]byte(``)),
			ExpectedCode: 400,
			ID:           "2",
		},
	}

	for _, test := range tests {
		pathWithId := fmt.Sprintf(test.Path, test.ID)
		req, err := http.NewRequest(test.Method, pathWithId, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")

		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		if res.Result().StatusCode != test.ExpectedCode {
			t.Fatalf("Expected Code: %v, Actual Code: %v.", test.ExpectedCode, res.Result().StatusCode)
		}
		assert.Equal(t, res.Result().StatusCode, test.ExpectedCode)
	}
}

func TestDeleteBook(t *testing.T) {
	Init()
	HandleRequest()

	tests := []TestWithID{
		{Method: "DELETE", Path: "/api/v1/books/%s", Body: nil, ExpectedCode: 200, ID: "1"},
		{Method: "DELETE", Path: "/api/v1/books/%s", Body: nil, ExpectedCode: 404, ID: "1"},
		{Method: "DELETE", Path: "/api/v1/books/%s", Body: nil, ExpectedCode: 404, ID: "6"},
	}
	for _, test := range tests {
		url := fmt.Sprintf(test.Path, test.ID)
		req, err := http.NewRequest(test.Method, url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedCode)
	}
}

func TestGetBook(t *testing.T) {
	Init()
	HandleRequest()

	tests := []TestWithID{
		{Method: "GET", Path: "/api/v1/books/%s", Body: nil, ExpectedCode: 200, ID: "1"},
		{Method: "GET", Path: "/api/v1/books/%s", Body: nil, ExpectedCode: 404, ID: "8"},
		{Method: "GET", Path: "/api/v1/books/%s", Body: nil, ExpectedCode: 200, ID: "3"},
	}
	for _, test := range tests {
		url := fmt.Sprintf(test.Path, test.ID)
		req, err := http.NewRequest(test.Method, url, test.Body)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedCode)
	}
}

func TestGetBookByIsbn(t *testing.T) {
	Init()
	HandleRequest()

	tests := []TestWithIsbn{
		{Method: "GET", Path: "/api/v1/books/isbn/%s", Body: nil, ExpectedCode: 200, Isbn: "14-14-14"},
		{Method: "GET", Path: "/api/v1/books/isbn/%s", Body: nil, ExpectedCode: 404, Isbn: "20-20-20"},
		{Method: "GET", Path: "/api/v1/books/isbn/%s", Body: nil, ExpectedCode: 404, Isbn: "33-33-33"},
	}
	for _, test := range tests {
		url := fmt.Sprintf(test.Path, test.Isbn)
		req, err := http.NewRequest(test.Method, url, test.Body)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedCode)
	}
}
