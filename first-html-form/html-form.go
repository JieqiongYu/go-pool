package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

func login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
	parsedTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
