package main

import (
	"net/http"
	"os"

	"github.com/go-pool/Go-Projects/Gopher-Face/endpoints"
	"github.com/go-pool/Go-Projects/Gopher-Face/handlers"
	"github.com/go-pool/Go-Projects/Gopher-Face/middleware"
	ghandlers "github.com/gorilla/handlers"
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

	// http://localhost:8080/restapi/socialmediapost/MerJQ
	r.HandleFunc("/restapi/socialmediapost/{username}", endpoints.FetchPostsEndpoint).Methods("GET")
	r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.CreatePostEndpoint).Methods("POST")
	r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.UpdatePostEndpoint).Methods("PUT")
	r.HandleFunc("/restapi/socialmediapost/{postid}", endpoints.DeletePostEndpoint).Methods("DELETE")

	// http.Handle("/", r)

	// middleware in gorilla toolkit
	// http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))

	// User self defined middleware
	// http.Handle("/", middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r)))
	http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(WEBSERVERPORT, nil)
}
