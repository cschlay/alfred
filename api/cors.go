package api

import (
	"net/http"
	"os"
)


func AddCorsHeaders(w http.ResponseWriter) {
	if os.Getenv("DEBUG") == "1" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "https://alfred-app.laych.dev")
	}
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
}
