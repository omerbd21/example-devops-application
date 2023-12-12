package mail

import (
	"crypto/tls"
	"os"
	"strconv"

	gomail "gopkg.in/mail.v2"
)

var senderMail = os.Getenv("ADMIN_EMAIL")
var senderPassword = os.Getenv("ADMIN_PASSWORD")
var subject = os.Getenv("SUBJECT")
var smtpHost = os.Getenv("SMTP_HOST")
var smtpPort = os.Getenv("SMTP_PORT")

// SendMail recieves an email address and content and sends the content to the address
func SendMail(email string, msg string) error {

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		// Handle the error if the conversion fails
		return nil
	}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", senderMail)

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", msg)

	// Settings for SMTP server
	d := gomail.NewDialer(smtpHost, port, senderMail, senderPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
