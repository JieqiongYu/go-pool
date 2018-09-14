// Http Server with basic authentication function
package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	connHost      = "localhost"
	connPort      = "8080"
	adminUser     = "admin"
	adminPassword = "admin"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(adminUser)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(adminPassword)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(connHost+":"+connPort, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
