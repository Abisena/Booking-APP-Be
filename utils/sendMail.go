package utils

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)


func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func SendMail(subject string, body string, to []string){
	email := os.Getenv("EMAIL_APP")
	password := os.Getenv("SANDI_APP")


	auth := smtp.PlainAuth(
		"",
		email,
		password,
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
