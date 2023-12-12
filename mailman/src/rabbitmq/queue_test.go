package rabbitmq

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

type consumeTestCase struct {
	queue string
	pass  bool
}

func TestConsume(t *testing.T) {
	testCases := []consumeTestCase{
		{os.Getenv("TEST_QUEUE"), true},
		{"notGood", false},
	}
	// Create a new logrus logger
	logger := logrus.New()

	// Set the logger to write JSON formatted logs
	logger.SetFormatter(&logrus.JSONFormatter{})
	rabbitUsername := os.Getenv("TEST_USERNAME")
	rabbitPassword := os.Getenv("TEST_PASSWORD")
	emails := os.Getenv("TEST_EMAIL_ADDRESSES")
	rabbitServer := os.Getenv("TEST_HOST")
	for _, tc := range testCases {
		credentials := RabbitCredenetials{
			RabbitUsername: rabbitUsername,
			RabbitPassword: rabbitPassword,
			Queue:          tc.queue,
			Vhost:          "",
			Emails:         emails,
			RabbitServer:   rabbitServer,
		}

		err := Consume(logger, credentials)
		if err != nil && tc.pass || err == nil && !tc.pass {
			t.Errorf("Consume test is not working.")
		}
	}
}
