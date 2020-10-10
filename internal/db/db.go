package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func getDbURL() string {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	return dbURL
}

func Connect(onConnect func(*pgx.Conn)) {
	dbURL := getDbURL()
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	onConnect(conn)
}

func ExecuteStatements(statements *[]string) {
	for _, statement := range *statements {
		ExecuteStatement(statement)
	}
}

func ExecuteStatement(statement string) {
	dbURL := getDbURL()
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	_, queryErr := conn.Query(context.Background(), statement)
	if queryErr != nil {
		fmt.Println(queryErr)
	}
}
