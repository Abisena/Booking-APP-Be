package utils

import (
	"fmt"
	"net/smtp"
)


func SendMail(subject string, body string, to []string){
	auth := smtp.PlainAuth(
		"",
		"abisena1123@gmail.com",
		"swojswzpimauiihl",
		"smtp.gmail.com",
	)

	// headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + body

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
