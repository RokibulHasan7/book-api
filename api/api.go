package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RokibulHasan7/book-api/db"

	"github.com/RokibulHasan7/book-api/auth"

	"github.com/RokibulHasan7/book-api/handler"

	"github.com/go-chi/chi"
)

var Router = chi.NewRouter()

func Init() {
	db.InitAuthor()
	db.InitBook()
	db.InitUser()
}

func HandleRequest() {
	//router.Use(middleware.RequestID)
	//router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)
	//router.Use(middleware.URLFormat)

	// Protected routes
	Router.Group(func(r chi.Router) {
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

		// Logout
		r.Post("/logout", auth.Logout)
	})

	// Public Routes
	Router.Group(func(rc chi.Router) {
		rc.Post("/login", auth.Login)

		rc.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Home."))
		})

		// Get AllBooks
		rc.Get("/api/v1/books", handler.GetAllBooks)

		// Get Book By Id
		rc.Get("/api/v1/books/{id}", handler.GetBook)

		// Get Book By ISBN
		rc.Get("/api/v1/books/isbn/{isbn}", handler.GetBookByIsbn)
	})

}

func StartServer(portNum string) {
	Init()          // Initialize DB
	HandleRequest() // Expose Routers

	// Server start
	sigs := make(chan os.Signal, 1) // Channel created to get the notification of Interrupt
	signal.Notify(sigs, os.Interrupt)
	portNum = ":" + portNum
	go func() {
		if err := http.ListenAndServe(portNum, Router); err != nil {
			log.Printf("Shutting down, reason: %s", err.Error())
			return
		}
	}()
	log.Printf("Server is listening on port %v", portNum)
	<-sigs

	time.Sleep(2 * time.Second)
	log.Println("Server is shutting down")
}
