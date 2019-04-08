package handlers

import "net/http"

/*MyProfileHandler is the handler for MyProfile*/
func MyProfileHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("profile"))
}
