package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type loginData struct {
	Username      string `json:username`
	Password      string `json:password`
}

func Login(w http.ResponseWriter, r *http.Request) {
	AddCorsHeaders(w)
	if r.Method == "POST" {
		var data loginData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(data.Password)
		fmt.Println(data.Username)
	} else if r.Method == "OPTIONS" {
		json.NewEncoder(w).Encode("OKOK")
	}
}