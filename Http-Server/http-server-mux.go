// Http Server with GZIP compression function
package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	connHost = "localhost"
	connPort = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(connHost+":"+connPort, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
