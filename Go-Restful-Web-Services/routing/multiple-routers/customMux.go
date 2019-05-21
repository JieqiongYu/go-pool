package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

/* CustomServeMux is a struct which can be a multiplexer*/
type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {
	// Any struct that has serveHTTP function can be a multiplexer
	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
