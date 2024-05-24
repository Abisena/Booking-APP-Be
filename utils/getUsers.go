package utils

import (
	"booking-app/auth"
	"database/sql"
	"fmt"
)

func GetUser(db *sql.DB) (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var password string
	var beliTiket uint

	fmt.Print("Masukan field dibawah ini\n")
	fmt.Println("Masukan Nama pertama:")
	fmt.Scan(&firstName)

	fmt.Println("Masukan Nama terakhir:")
	fmt.Scan(&lastName)

	fmt.Println("Masukan Alamat email:")
	fmt.Scan(&email)
	
	fmt.Println("Masukan Alamat password:")
	fmt.Scan(&password)

	fmt.Println("Masukan Jumlah tiket:")
	fmt.Scan(&beliTiket)

	auth.Register(db, firstName, email, password)

	return firstName, lastName, email, beliTiket
}


