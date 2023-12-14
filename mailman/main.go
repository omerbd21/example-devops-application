package main

import (
	"os"
	"strings"
	"sync"

	"github.com/omerbd21/mailman/src/mail"
	"github.com/omerbd21/mailman/src/rabbitmq"
	"github.com/sirupsen/logrus"
)

var (
	rabbitUsername string
	rabbitPassword string
	queue          string
	vhost          string
	emails         string
	rabbitServer   string
	port           string
)

func init() {
	rabbitUsername = os.Getenv("RABBIT_USERNAME")
	rabbitPassword = os.Getenv("RABBIT_PASSWORD")
	queue = os.Getenv("QUEUE")
	vhost = os.Getenv("VIRTUAL_HOST")
	emails = os.Getenv("EMAIL_ADDRESSES")
	rabbitServer = os.Getenv("RABBIT_HOST")
	port = os.Getenv("RABBIT_PORT")
}

func main() {
	// Create a new logrus logger
	logger := logrus.New()

	// Set the logger to write JSON formatted logs
	logger.SetFormatter(&logrus.JSONFormatter{})

	var wg sync.WaitGroup
	c := make(chan string)
	wg.Add(2)

	config := rabbitmq.RabbitConfig{
		Username:  rabbitUsername,
		Password:  rabbitPassword,
		Host:      rabbitServer,
		Port:      port,
		VHost:     vhost,
		QueueName: queue,
	}
	r := rabbitmq.Rabbit{Config: config, Log: logger}
	err := r.Connect()
	if err != nil {
		logger.WithError(err)
	}

	err = r.Channel()
	if err != nil {
		logger.WithError(err)
	}

	emailArray := strings.Split(emails, ",")

	go r.Consume(c, &wg)
	go mail.SendStories(c, &wg, emailArray)
	wg.Wait()

}
