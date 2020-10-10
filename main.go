package main

import (
    "fmt"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
}

func main() {
    handleRequests()
    fmt.Println("Server up at localhost:8000.")
    http.ListenAndServe(":8000", nil)
}
