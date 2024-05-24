package utils

import (
	"fmt"
	"net/smtp"
)


func SendMail(subject string, html string, to []string){
	auth := smtp.PlainAuth(
		"",
		"abisena1123@gmail.com",
		"231222",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject" + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"abisena1123@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
