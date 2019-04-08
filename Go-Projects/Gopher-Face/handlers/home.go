package handlers

import (
	"fmt"
	"net/http"
)

/*HomeHandler is the handler for Home*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
