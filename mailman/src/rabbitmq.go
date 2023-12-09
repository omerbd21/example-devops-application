package src

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Consume(logger *logrus.Logger) {
	rabbitUsername := os.Getenv("USERNAME")
	rabbitPassword := os.Getenv("PASSWORD")
	queue := os.Getenv("QUEUE")
	vhost := os.Getenv("VIRTUAL_HOST")
	emails := os.Getenv("EMAIL_ADDRESSES")
	rabbitServer := os.Getenv("HOST")
	connection, err := amqp.Dial("amqp://" + rabbitUsername + ":" + rabbitPassword + "@" + rabbitServer + ":5672/")
	if err != nil {
		logger.WithError(err).Error("Failed to connect to rabbitmq")
	}
	defer connection.Close()

	ampqChannel, err := connection.Channel()
	if err != nil {
		logger.WithError(err).Error("Failed to create a channel")
	}
	defer ampqChannel.Close()

	messageChannel, err := ampqChannel.Consume(
		queue,
		vhost,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.WithError(err).Error("Failed to consume messages")
	}

	stopChan := make(chan bool)
	go func() {
		for d := range messageChannel {
			emailAddresses := strings.Split(emails, ",")
			for _, email := range emailAddresses {
				msg := string(d.Body[:])
				err := SendMail(email, msg)
				logger.WithFields(logrus.Fields{"characters": len(msg), "email": email}).Info("Sent a mail")
				if err != nil {
					logger.WithError(err).Error("Failed to send a mail")
				}
			}
			if err := d.Ack(false); err != nil {
				logger.WithError(err).Error("Failed to acknowledge the message")
			}
		}
	}()

	<-stopChan

}
