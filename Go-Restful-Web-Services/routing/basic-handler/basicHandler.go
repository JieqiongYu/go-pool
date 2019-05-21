package main

import (
	"io"
	"log"
	"net/http"
)

/*MyServe is the web server for saying hello world*/
// Here `w` is a response writer. A `ResponseWriter` interface is used by an HTTP handler to construct an HTTP response.
// `req` is a request object, which deals with all the properties and methods of an HTTP request.
func MyServe(w http.ResponseWriter, req *http.Request) {
	// 4. Write `hello, world!` response.
	io.WriteString(w, "hello, world!\n")
}

func main() {
	// 1. Create a route called /hello
	// 2. Create a handler called MyServe
	// 3. Whenever the request comes on the route (/hello), the handler function will be executed.
	// The `http` package has a function called `HandleFunc`, using which we can map an URL to a function.
	http.HandleFunc("/hello", MyServe)
	// 5. Start the server on port 8000. `ListenAndServe` returns `error` if somethings goes wrong. So log it using `log.Fatal`.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
