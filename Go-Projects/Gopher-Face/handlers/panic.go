package handlers

import "net/http"

/*TriggerPanicHandler is the handler for Trigger Panic*/
func TriggerPanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Triggering a Panic!")
}
