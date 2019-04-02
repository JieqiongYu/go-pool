package main

import (
	"fmt"
	"net/http"

	validationkit "github.com/go-pool/Go-Basics/validation-kit"
)

func main() {
	// curl http://localhost:8080/hello-gopher
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	// curl http://localhost:8080/username-syntax-checker?username=MerJQ
	http.HandleFunc("/username-syntax-checker", checkUsernameSyntaxHandler)
	http.ListenAndServe(":8080", nil)
}

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}

func checkUsernameSyntaxHandler(w http.ResponseWriter, r *http.Request) {
	var usernameSyntaxResult bool
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username not provided!", http.StatusInternalServerError)
	} else {
		usernameSyntaxResult = validationkit.CheckUsernameSyntax(username)
		fmt.Fprintf(w, "Syntax Check Result for %v is %v", username, usernameSyntaxResult)
	}
}
