package sendtiket

import (
	"booking-app/utils"
	"fmt"
	"sync"
	"time"
)


var wg = sync.WaitGroup{}

func SendTiketUser(firstName string, lastName string, email string, beliTiket uint) {
	time.Sleep(5 *time.Second)
	var tiket = fmt.Sprintf("%v tikets akan diberikan ke %v %v", beliTiket, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", tiket, email)
	utils.SendMail("Booking-app", "Thank You for Booking app in this web", []string{email})
	fmt.Println("#################")
	wg.Done()
}