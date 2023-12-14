package mail

import (
	"crypto/tls"
	"strconv"
	"sync"

	gomail "gopkg.in/mail.v2"
)

// SendMail recieves an email address and content and sends the content to the address
func SendMail(email string, msg string) error {

	port, err := strconv.Atoi(SmtpPort)
	if err != nil {
		// Handle the error if the conversion fails
		return nil
	}

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", Address)

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("Subject", Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", msg)

	// Settings for SMTP server
	d := gomail.NewDialer(SmtpHost, port, Address, Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// SendStories gets a string Channel, a pointer to a WaitGroup and a slice of email addresses
// and uses SendMail in order to send the stories from the channel to the email addresses.
func SendStories(c chan string, wg *sync.WaitGroup, emails []string) {
	defer wg.Done()
	for msg := range c {
		for _, email := range emails {
			SendMail(email, msg)
		}
	}
}
