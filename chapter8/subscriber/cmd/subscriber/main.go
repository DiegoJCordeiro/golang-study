package main

import (
	"fmt"
	"github.com/DiegoJCordeiro/golang-study/chapter8/subscriber/pkg/rabbitmq"
)

func main() {

	rabbitMQ := rabbitmq.NewRabbitMQ()
	err := rabbitMQ.OpenConnection()

	if err != nil {
		panic(err)
	}

	go func() {
		err := rabbitMQ.Consume("queue-rabbitmq-golang")
		if err != nil {
			panic(err)
		}
	}()

	for message := range rabbitMQ.GetChannelDelivery() {

		fmt.Println(string(message.Body))
		_ = message.Ack(false)
	}
}
