package main

import (
	"gopkg.in/gomail.v2"
)

type UserInput struct {
	Email   string
	Subject string
	Message string
}

func main() {
	input := UserInput{"example email", "example subject", "example message"}
	generateAndSendEmail(input)
}

func generateAndSendEmail(input UserInput) {
	m := gomail.NewMessage()
	m.SetHeader("From", "Aristos.Website@gmail.com")
	m.SetHeader("To", "aristos.a.athens@gmail.com")
	m.SetHeader("Subject", input.Subject)
	m.SetBody("text/html", input.Message)

	d := gomail.NewDialer("smtp.gmail.com", 587, "Aristos.Website", "VerySecurePassword")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
