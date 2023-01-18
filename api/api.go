package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RokibulHasan7/book-api/db"

	"github.com/go-chi/chi/middleware"

	"github.com/RokibulHasan7/book-api/auth"

	"github.com/RokibulHasan7/book-api/handler"

	"github.com/go-chi/chi"
)

func init() {
	os.Setenv("ADMIN_USER", "admin")
	os.Setenv("ADMIN_PASS", "1234")
	db.InitAuthor()
	db.InitBook()
	db.InitUser()
}

func HandleRequest() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// Protected routes
	router.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		//r.Use(jwtauth.Verifier(db.TokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		//r.Use(jwtauth.Authenticator)

		r.Use(auth.PrimaryAuth)

		// Post Book
		r.Post("/api/v1/books", handler.PostBook)

		// Update Book
		r.Put("/api/v1/books/{id}", handler.UpdateBook)

		// Delete Book By Id
		r.Delete("/api/v1/books/{id}", handler.DeleteBook)
	})

	// Public Routes
	router.Group(func(rc chi.Router) {
		rc.Post("/login", auth.Login)
		rc.Post("/logout", auth.Logout)
		rc.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Home."))
		})

		// Get AllBooks
		rc.Get("/api/v1/books", handler.GetAllBooks)

		// Get Book By Id
		rc.Get("/api/v1/books/{id}", handler.GetBook)
	})

	// Server start
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		if err := http.ListenAndServe(":3333", router); err != nil {
			log.Println("Shutting down")
			log.Fatalln(err)
			return
		}
	}()
	log.Println("Server is listening on port 3333")
	<-sigs
	time.Sleep(2 * time.Second)
	log.Println("Server is shutting down")

}
