package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=admin12345 host=localhost port=5432")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully Connected")
	return db, nil
}