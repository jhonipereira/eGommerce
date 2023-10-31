package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/jhonipereira/go-micro-listener/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitMQ
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	// start listening for messages
	log.Println("listening for and consuming RabbitMQ messages... ")

	// create the consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Panic(err)
	}

	// watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't procceed until rabbit is UP
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("rabbitmq not yet ready...")
			counts++
		} else {
			log.Println("connect to rabbitmq")
			connection = c
			break
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off")
		time.Sleep(backOff)
	}

	return connection, nil

}
