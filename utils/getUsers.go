package utils

import (
	"fmt"
)

func GetUser() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var beliTiket uint

	fmt.Print("Masukan field dibawah ini\n")
	fmt.Println("Masukan Nama pertama:")
	fmt.Scan(&firstName)

	fmt.Println("Masukan Nama terakhir:")
	fmt.Scan(&lastName)

	fmt.Println("Masukan Alamat email:")
	fmt.Scan(&email)

	fmt.Println("Masukan Jumlah tiket:")
	fmt.Scan(&beliTiket)

	return firstName, lastName, email, beliTiket
}


