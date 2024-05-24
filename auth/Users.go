package auth

import (
	"database/sql"
	"fmt"
	"log"
)


type UserModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(db *sql.DB, email string, password string) (string, string) {
	fmt.Println("Silahkan Register")

	stmt, err := db.Prepare("INSERT INTO users (email, password) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}

	user := UserModel{
		Email:    email,
		Password: password,
	}

	_, err = stmt.Exec(user)
	if err != nil {
		log.Fatal(err)
	}

	return email, password
}
