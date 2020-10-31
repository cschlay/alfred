package api

import (
	"alfred/internal/accounts"
	"encoding/json"
	"fmt"
	"net/http"
)

type NewUserData struct {
	Username      string `json:username`
	Password      string `json:password`
	PasswordAgain string `json:passowdAgain`
}

func Users(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data NewUserData

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if data.Password != data.PasswordAgain {
			fmt.Println("NOT SAME")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(data)
			return
		}
		accounts.CreateAccount(data.Username, data.Password)
	}
}
