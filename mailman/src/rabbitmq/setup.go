package rabbitmq

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	VHost     string
	QueueName string
}

type Rabbit struct {
	Config     RabbitConfig
	connection *amqp.Connection
	chn        *amqp.Channel
	Log        *logrus.Logger
}

// Connect uses the Rabbit struct in order to connect to a RabbitMQ instance.
func (r *Rabbit) Connect() error {
	if r.connection == nil || r.connection.IsClosed() {
		conn, err := amqp.Dial(fmt.Sprintf(
			"amqp://%s:%s@%s:%s/%s",
			r.Config.Username,
			r.Config.Password,
			r.Config.Host,
			r.Config.Port,
			r.Config.VHost,
		))
		if err != nil {
			r.Log.WithError(err).Error("Failed to connect to RabbitMQ")
			return err
		}
		r.connection = conn
	}

	return nil
}

// Channel creates a new RabbitMQ channel and stores it into the chn property of the Rabbit struct
func (r *Rabbit) Channel() error {
	channel, err := r.connection.Channel()
	if err != nil {
		r.Log.WithError(err).Error("Failed to create a channel.")
		return err
	}
	r.chn = channel

	return err
}
