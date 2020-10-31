package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"

	"alfred/api"
	"alfred/internal/db"
)

func handleRequests() {
	http.HandleFunc("/z0I1tzHowf0OkaKR/users", api.Users)
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
	db.Connect(testDb)
	db.SyncTables()

	handleRequests()
	fmt.Println("Server up at localhost:8000.")
	http.ListenAndServe(":8000", nil)
}
