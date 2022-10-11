package core

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail(recipient, subject, message string) bool {
	email := gomail.NewMessage()

	email.SetHeader("From", "louis2001saad@gmail.com")
	email.SetHeader("To", recipient)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", message)

	mail := gomail.NewDialer("smtp.gmail.com", 587, "louistestodoo@gmail.com", "ixyvklsygvfktqeg")

	if err := mail.DialAndSend(email); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
