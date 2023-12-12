package rabbitmq

import (
	"strings"
	"sync"

	"github.com/omerbd21/mailman/src/mail"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitCredenetials struct {
	RabbitUsername string
	RabbitPassword string
	RabbitServer   string
	Queue          string
	Vhost          string
	Emails         string
}

func Consume(logger *logrus.Logger, credentials RabbitCredenetials) error {
	connection, err := amqp.Dial("amqp://" + credentials.RabbitUsername + ":" + credentials.RabbitPassword + "@" + credentials.RabbitServer + ":5672/")
	if err != nil {
		logger.WithError(err).Error("Failed to connect to rabbitmq")
		return err
	}
	defer connection.Close()

	ampqChannel, err := connection.Channel()
	if err != nil {
		logger.WithError(err).Error("Failed to create a channel")
		return err
	}
	defer ampqChannel.Close()

	messageChannel, err := ampqChannel.Consume(
		credentials.Queue,
		credentials.Vhost,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.WithError(err).Error("Failed to consume messages")
		return err
	}

	var wg sync.WaitGroup
	var consumeErr error
	stopChan := make(chan bool)
	go func() {
		wg.Add(1)
		defer wg.Done()
		for d := range messageChannel {
			emailAddresses := strings.Split(credentials.Emails, ",")
			for _, email := range emailAddresses {
				msg := string(d.Body[:])
				err := mail.SendMail(email, msg)
				logger.WithFields(logrus.Fields{"characters": len(msg), "email": email}).Info("Sent a mail")
				if err != nil {
					logger.WithError(err).Error("Failed to send a mail")
					consumeErr = err
					return
				}
			}
			if err := d.Ack(false); err != nil {
				logger.WithError(err).Error("Failed to acknowledge the message")
				consumeErr = err
				return
			}
		}
	}()

	go func() {
		// Wait for the goroutine to finish
		wg.Wait()

		// Signal that the goroutine has finished
		stopChan <- true
		close(stopChan)
	}()

	// Wait for the goroutine to finish or an error to occur
	select {
	case <-stopChan:
		// Return the error encountered in the goroutine, if any
		return consumeErr
	}

}
