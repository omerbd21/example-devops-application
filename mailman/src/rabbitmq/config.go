package rabbitmq

import "sync"

// queueDeclare uses the Rabbit struct in order to declare a queue
func (r *Rabbit) queueDeclare() error {
	if _, err := r.chn.QueueDeclare(
		r.Config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}
	return nil
}

// Consume uses the Rabbit struct and gets a string Channel and a pointer to a WaitGroup
// and consumes the messages into the channel
func (r *Rabbit) Consume(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msgs, err := r.chn.Consume(
		r.Config.QueueName,
		r.Config.VHost,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.Log.WithError(err).Error("Consumer failed to start")
		return
	}

	r.Log.Info("Mailman running ...")

	for msg := range msgs {
		r.Log.Info("Consumed:", string(msg.Body))
		c <- string(msg.Body)
		if err := msg.Ack(false); err != nil {
			r.Log.WithError(err).Error("Acknowledge failed, dropped message", err)
		}
	}
	close(c)

}
