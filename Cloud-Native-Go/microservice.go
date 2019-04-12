package main

import (
	"fmt"
	"net/http"
	"os"

	api "github.com/go-pool/Cloud-Native-Go/api"
)

func main() {
	// http://localhost:8080/
	http.HandleFunc("/", index)
	// localhost:8080/api/echo?message=Cloud-Native-Go
	http.HandleFunc("/api/echo", echo)

	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello Cloud Native Go.")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
