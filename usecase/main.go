package main

import (
	"encoding/json"
	"go-smtp-service/Core/mail"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//m := gomail.NewMessage()
	//m.SetHeader("From", os.Getenv("FORCE_TO"))
	//m.SetHeader("To", os.Getenv("FORCE_TO"))
	//m.SetHeader("Subject", "Hello! From Go ")
	//m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	payload := []byte("sdasdasdsa")

	port, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))
	mailer := mail.NewMailer(os.Getenv("MAILER_HOST"), os.Getenv("MAILER_USERNAME"), os.Getenv("MAILER_PASSWORD"), port)

	m := mail.NewMessage()
	err = json.Unmarshal(payload, &m)

	// Send the email to Bob, Cora and Dan.
	if err := mailer.Send(m); err != nil {
		panic(err)
	}
}
