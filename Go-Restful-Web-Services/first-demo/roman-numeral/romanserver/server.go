package main

import (
	"fmt"
	"html"

	// Core package we used to handle an HTTP request through its `HandleFunc` function.
	// That function's arguments are `http.Request` and `http.ResponseWriter`.
	// Those two deal with the request and response of an HTTP request.
	"net/http"
	"strconv"
	"strings"

	// The `time` package is used to define seconds in the program.
	"time"

	// Here is the data service we created before.
	romannumerals "github.com/restful-web-services/roman-numeral/romanNumerals"
)

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// r.URL.Path is the URL path of the HTTP request.
		// For the CURL Request one, it is /roman_number/5.
		// We are splitting this path and using the second argument as a resource
		// and the third argument as a value to get the Roman numeral.
		// The `Split` function is in a core package called `strings`.
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			// The `Atoi` function converts an alphanumeric string to an integer.
			// For the numerals map to consume, we need to convert the integer string to an integer.
			// The `Atoi` function comes from a core package called `strconv`.
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is not in the list, send NOt Found status
				// We use `http.StatusXXX` to set the status code of the response header.
				// The `WriteHeader` and `Write` functions are available on the response object
				// for writing the header and body, respectively.
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				// `EscapeString` escapes special characters to become valid HTML characters.
				// For example, `Fran & Freddie's` becomes `Fran &amp; Freddie's&#34`
				fmt.Fprintf(w, "%q", html.EscapeString(romannumerals.Numerals[number]))
			}
		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}
	})

	// Create a server and run it on 8000 port
	// We created an HTTP server using `&http` while initializing a few parameters like
	// address, port, timeout, and so on.
	s := &http.Server{
		Addr: ":8000",
		// It says, after 10 seconds of inactivity, automatically return a 408 request timeout back to the client.
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Finally, start the server with the `ListenAndServe` function.
	// It keeps your web server running until you kill it.
	s.ListenAndServe()
}
