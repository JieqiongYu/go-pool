package main

import (
	"net/http"

	"github.com/go-pool/Go-Projects/Gopher-Face/handlers"
	mux "github.com/gorilla/mux"
)

/**WEBSERVERPORT*/
const (
	WEBSERVERPORT = ":8080"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/triggerpanic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")

	// r.HandleFunc("/restapi/socialmediapost/{username}", endpoints.FetchPostsEndpoint).Methods("GET")
	// r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
	// r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	// r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")

	http.Handle("/", r)

	http.ListenAndServe(WEBSERVERPORT, nil)
}
