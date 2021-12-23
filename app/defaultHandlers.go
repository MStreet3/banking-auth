package app

import "net/http"

func pong(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, "pong")
}
