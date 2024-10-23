package rabbitmq

import "github.com/streadway/amqp"

type Rabbitmq struct {
	channel         *amqp.Channel
	connection      *amqp.Connection
	channelDelivery chan amqp.Delivery
}

func (rabbitmq *Rabbitmq) GetChannel() *amqp.Channel {
	return rabbitmq.channel
}

func (rabbitmq *Rabbitmq) GetConnection() *amqp.Connection {
	return rabbitmq.connection
}

func (rabbitmq *Rabbitmq) GetChannelDelivery() chan amqp.Delivery {
	return rabbitmq.channelDelivery
}

func NewRabbitMQ() *Rabbitmq {
	return &Rabbitmq{}
}

func (rabbitmq *Rabbitmq) OpenConnection() error {

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		return err
	}
	rabbitmq.connection = connection

	channelConn, err := connection.Channel()

	if err != nil {
		return err
	}

	rabbitmq.channel = channelConn
	rabbitmq.channelDelivery = make(chan amqp.Delivery)

	return nil
}

func (rabbitmq *Rabbitmq) Consume(queueName string) error {
	messages, err := rabbitmq.GetChannel().Consume(
		queueName,
		"golang-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for message := range messages {
		rabbitmq.channelDelivery <- message
	}

	return nil
}

func (rabbitmq *Rabbitmq) Publish(payload, routingKey, exchangeName string) error {

	errPublish := rabbitmq.GetChannel().Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	return errPublish
}
