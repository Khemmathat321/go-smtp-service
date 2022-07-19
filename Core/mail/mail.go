package mail

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"gopkg.in/gomail.v2"
)

type Attachment struct {
	Url      string
	Filename string
	File     string
	Mimetype string
}

type Message struct {
	CredentialKey string
	To            string
	FromEmail     string
	FromName      string
	Body          string
	Subject       string
	Attachments   []Attachment
}

type Mailer struct {
	connection *gomail.Dialer
}

func NewMailer(host, username, password string, port int) *Mailer {
	m := &Mailer{}
	m.Connect(
		host,
		username,
		password,
		port,
	)

	return m
}

func NewMessage() *Message {
	return new(Message)
}

func (m *Mailer) Connect(host, username, password string, port int) {
	m.connection = gomail.NewDialer(
		host,
		port,
		username,
		password,
	)
}

func (m *Mailer) Send(msg *Message) error {
	message := gomail.NewMessage()

	if f := os.Getenv("FORCE_TO"); f != "" {
		msg.To = f
		msg.Subject = "FORCED: " + msg.Subject
	}

	message.SetAddressHeader("From", msg.FromEmail, msg.FromName)
	message.SetHeader("To", msg.To)
	message.SetHeader("Subject", msg.Subject)
	message.SetBody("text/html", msg.Body)

	for _, att := range msg.Attachments {
		message.Attach(att.Filename, gomail.SetCopyFunc(func(w io.Writer) error {
			client := http.Client{}
			resp, _ := client.Get(att.Url)

			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			_, err := w.Write(buf.Bytes())

			return err
		}))
	}

	return m.connection.DialAndSend(message)
}
