package helper

import (
	"strings"
)

var sisaTiket uint = 250
func ValidasiData(firstName string, lastName string, email string, beliTiket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := beliTiket > 0 && beliTiket <= sisaTiket
	return isValidName, isValidEmail, isValidTicketNumber
}