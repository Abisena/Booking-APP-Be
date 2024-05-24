package auth

import (
	"database/sql"
	"fmt"
	"log"
)


type UserModel struct {
	ID       int `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *sql.DB, username string, email string, password string) (string, string, string) {
	fmt.Println("Silahkan Register")

	stmt, err := db.Prepare("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}

	user := UserModel{
		Username: username,
		Email:    email,
		Password: password,
	}

	_, err = stmt.Exec(user)
	if err != nil {
		log.Fatal(err)
	}

	return username, email, password
}
