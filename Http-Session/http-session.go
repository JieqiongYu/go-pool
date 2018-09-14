package main

import (
	"github.com/gorilla/sessions"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func main() {
	// http.HandleFunc("/home", home)
	// http.HandleFunc("login", login)
	// http.HandleFunc("logout", logout)
}
