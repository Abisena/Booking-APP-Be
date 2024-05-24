package main

import (
	"booking-app/database"
	"booking-app/helper"
	"booking-app/sendtiket"
	"booking-app/transaction"
	"booking-app/utils"
	"fmt"
	"log"
	"sync"
)


var wg = sync.WaitGroup{}

var sisaTiket uint = 250

func main() {
	// Konecksi Database
	db, err := database.Connect()
    if err != nil {
      log.Fatal(err)
      return
    }
    defer db.Close()
    // Convert *sqlx.DB to *sql.DB
    // sqlDB := db.DB
	// End Close

	fmt.Println("########### Welcome to Booking App ###########")

	firstName, lastName, email, beliTiket := utils.GetUser()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidasiData(firstName, lastName, email, beliTiket)
	

	if isValidName && isValidEmail && isValidTicketNumber {
		transaction.TransactionBooking(firstName, lastName, email, beliTiket)
	

		wg.Add(1)
		go sendtiket.SendTiketUser(firstName, lastName, email, beliTiket)
		
		if sisaTiket == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}
