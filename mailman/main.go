package main

import (
	"os"

	"github.com/omerbd21/mailman/src/rabbitmq"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new logrus logger
	logger := logrus.New()

	// Set the logger to write JSON formatted logs
	logger.SetFormatter(&logrus.JSONFormatter{})

	rabbitUsername := os.Getenv("USERNAME")
	rabbitPassword := os.Getenv("PASSWORD")
	queue := os.Getenv("QUEUE")
	vhost := os.Getenv("VIRTUAL_HOST")
	emails := os.Getenv("EMAIL_ADDRESSES")
	rabbitServer := os.Getenv("HOST")
	credentials := rabbitmq.RabbitCredenetials{
		RabbitUsername: rabbitUsername,
		RabbitPassword: rabbitPassword,
		Queue:          queue,
		Vhost:          vhost,
		Emails:         emails,
		RabbitServer:   rabbitServer,
	}

	rabbitmq.Consume(logger, credentials)
}
