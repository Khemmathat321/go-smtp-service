package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
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

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("FORCE_TO"))
	m.SetHeader("To", os.Getenv("FORCE_TO"))
	m.SetHeader("Subject", "Hello! From Go ")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	port, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))
	d := gomail.NewDialer(os.Getenv("MAILER_HOST"), port, os.Getenv("MAILER_USERNAME"), os.Getenv("MAILER_PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
