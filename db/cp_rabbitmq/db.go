package cp_rabbitmq

import (
	"app-container-platform/config"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var rmq *RabbitMQClient

func InitRabbitMQConnection() error {
	rmq, err := connectToRabbitMQ(config.RabbitMQConnectionUrl)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer rmq.Connection.Close()
	defer rmq.Channel.Close()
	fmt.Println("Connected to RabbitMQ!")
	return nil
}

func connectToRabbitMQ(url string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("error connecting to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("error opening RabbitMQ channel: %v", err)
	}

	return &RabbitMQClient{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func PublishToRabbitMQ(queueName, message string) error {
	err := rmq.Channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return errors.New("error publishing message to RabbitMQ: %v")
	}
	return nil
}

func ConsumeFromRabbitMQ(queueName string) (<-chan amqp.Delivery, error) {
	log.Println("consume 1")
	msgs, err := rmq.Channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, fmt.Errorf("error consuming messages from RabbitMQ: %v", err)
	}
	return msgs, nil
}
