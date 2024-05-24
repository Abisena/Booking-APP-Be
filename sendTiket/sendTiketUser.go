package sendtiket

import (
	"fmt"
	"sync"
	"time"
)


var wg = sync.WaitGroup{}

func SendTiketUser(firstName string, lastName string, email string, beliTiket uint) {
	time.Sleep(20 *time.Second)
	var tiket = fmt.Sprintf("%v tikets akan diberikan ke %v %v", beliTiket, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", tiket, email)
	fmt.Println("#################")
	wg.Done()
}