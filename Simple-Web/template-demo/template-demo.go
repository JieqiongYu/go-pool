// Simple example of creating and using a template in Go.
package main

import (
	"log"
	"net/http"
	"text/template"
)

/*Gopher is a struct*/
// Type representing a gopher
type Gopher struct {
	Name string
}

func main() {
	// http://localhost:8080/hello-gopher?gophername=MerJQ
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.ListenAndServe(":8080", nil)
}

// Handler for the hello-gopher route
func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Gopher"
	}
	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./templates/greeting.html", gopher)
}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing and template:", err)
	}
	t.Execute(w, templateData)
}
