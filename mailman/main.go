package main

import (
	"github.com/omerbd21/mailman/src"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new logrus logger
	logger := logrus.New()

	// Set the logger to write JSON formatted logs
	logger.SetFormatter(&logrus.JSONFormatter{})
	src.Consume(logger)
}
