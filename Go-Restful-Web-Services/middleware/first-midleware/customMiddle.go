package main

import (
	"fmt"
	"net/http"
)

func main() {
	// HnalderFunc returns a HTTP Handler
	// Create a handler function by passing the main handler function (mainLogic) to http.HandlerFunc().
	mainLogicHandlder := http.HandlerFunc(mainLogic)
	// The http.Handle function expects an HTTP handler. By taking that into consideration, we wrapped our logic
	// in such a way that, finally, a handler get returned, but the execution is modified.
	http.Handle("/", middleware(mainLogicHandlder))
	http.ListenAndServe(":8000", nil)
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}

// Create a middleware function that accepts a handler and returns a handler
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		// The method ServeHTTP allows a handler to execute the handler logic that is mainLogic.
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}
