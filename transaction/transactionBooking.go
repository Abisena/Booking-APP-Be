package transaction

import "fmt"

const jumlahTiket uint = 250

var sisaTiket uint = 250
var Bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func TransactionBooking(firstName string, lastName string, email string, beliTiket uint) {
	sisaTiket := sisaTiket - beliTiket

	var DataUser = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: sisaTiket,
	}

	bookings := append(Bookings, DataUser)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, beliTiket, email)
}