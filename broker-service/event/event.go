package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // type
		true,         // durable ?
		false,        // auto delete?
		false,        // internal ?
		false,        // noWait ?
		nil,          // args
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable ?
		false, // auto delete ? when unused?
		true,  // exclusive?
		false, // no wait?
		nil,   // args
	)
}
