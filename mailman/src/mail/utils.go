package mail

import (
	"os"
)

var (
	Address  string
	Password string
	Subject  string
	SmtpHost string
	SmtpPort string
)

func init() {
	Address = os.Getenv("ADMIN_EMAIL")
	Password = os.Getenv("ADMIN_PASSWORD")
	Subject = os.Getenv("SUBJECT")
	SmtpHost = os.Getenv("SMTP_HOST")
	SmtpPort = os.Getenv("SMTP_PORT")
}
