package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"

	"alfred/internal/db"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
}

func testDb(conn *pgx.Conn) bool {
	fmt.Println("Database connection OK.")
	return true
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file missing, put it to the same directroy as main.")
		os.Exit(1)
	}

	handleRequests()
	db.Connect(testDb)
	db.SyncTables()

	fmt.Println("Server up at localhost:8000.")
	http.ListenAndServe(":8000", nil)
}
