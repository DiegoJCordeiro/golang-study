package main

import "github.com/DiegoJCordeiro/golang-study/chapter8/subscriber/pkg/rabbitmq"

func main() {

	rabbitMQ := rabbitmq.NewRabbitMQ()
	errOpenConnection := rabbitMQ.OpenConnection()

	if errOpenConnection != nil {
		panic(errOpenConnection)
	}

	errPublish := rabbitMQ.Publish("Hello World, RabbitMQ", "", "exchange-rabbitmq-golang")

	if errPublish != nil {
		panic(errOpenConnection)
	}
}
