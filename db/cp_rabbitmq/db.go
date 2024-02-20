package cp_rabbitmq

import (
	"app-container-platform/config"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//type RabbitMQClient struct {
//	Connection *amqp.Connection
//	Channel    *amqp.Channel
//}
//
//var rmq *RabbitMQClient

// var rmqConnection *amqp.Connection
var rmqChannel *amqp.Channel

func InitRabbitMQConnection() error {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(config.RabbitMQConnectionUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	//defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	//defer ch.Close()
	rmqChannel = ch

	log.Println("Connected to RabbitMQ!")
	return nil
}

func PublishToRabbitMQ(queueName, message string) error {
	err := rmqChannel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Println("unable to publish : " + err.Error())
		return err
	}
	return nil
}

func ConsumeFromRabbitMQ(queueName string) (<-chan amqp.Delivery, error) {
	log.Println("consume 1")
	msgs, err := rmqChannel.Consume(
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
