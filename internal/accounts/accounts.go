package accounts

/* https://gowebexamples.com/password-hashing/ */

import (
	"alfred/internal/db"
	"context"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(username string, password string) {
	passwordHash, _ := hashPassword(password)

	db.Connect(func(conn *pgx.Conn) bool {
		_, err := conn.Exec(context.Background(), "INSERT INTO account (username, password) VALUES ($1, $2)", username, passwordHash)
		if err != nil {
			return false
		}
		return true
	})
}

func CheckPassword(username string, password string) bool {
	var hash *string
	db.Connect(func(conn *pgx.Conn) bool {
		err := conn.QueryRow(context.Background(), "SELECT (password) FROM account WHERE username=$1", username).Scan(&hash)
		if err != nil {
			return false
		}
		return true
	})
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
