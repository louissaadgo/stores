package emailing

import "gopkg.in/gomail.v2"

func SendEmail(recipient, subject, message string) bool {
	email := gomail.NewMessage()

	email.SetHeader("From", "louistestodoo@gmail.com")
	email.SetHeader("To", recipient)
	email.SetHeader("Subject", subject)
	email.SetBody("text/plain", message)

	mail := gomail.NewDialer("smtp.gmail.com", 587, "louistestodoo@gmail.com", "pmrtujqqcgoybbex")

	if err := mail.DialAndSend(email); err != nil {
		return false
	}

	return true
}
