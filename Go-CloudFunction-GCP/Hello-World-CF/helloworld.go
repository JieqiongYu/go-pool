package helloworld

import (
	"fmt"
	"net/http"
)

/*HelloWorld is the hello world function for cloud function*/
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
