package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8000", loggedRouter)
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}
