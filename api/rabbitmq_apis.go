package api

import (
	"app-container-platform/db/cp_rabbitmq"
	_type "app-container-platform/types"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type RabbitMQApiInf interface {
	ApiIndex(c echo.Context) error
	Publish(c echo.Context) error
	Consume(c echo.Context) error
}

type rabbitmqApi struct{}

func (r rabbitmqApi) ApiIndex(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "App Container Platform Rabbitmq Api",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (r rabbitmqApi) Publish(c echo.Context) error {
	queueName := c.Param("queue")
	message := c.QueryParam("message")
	if message == "" {
		return c.JSON(http.StatusBadRequest, "Message parameter is required")
	}
	err := cp_rabbitmq.PublishToRabbitMQ(queueName, message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error publishing message to RabbitMQ: %v")
	}
	return c.JSON(http.StatusOK, "Message published to RabbitMQ!")
}

func (r rabbitmqApi) Consume(c echo.Context) error {
	queueName := c.QueryParam("queue")
	msgs, err := cp_rabbitmq.ConsumeFromRabbitMQ(queueName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error consuming messages from RabbitMQ: %v")
	}
	for msg := range msgs {
		log.Println("Message: ", string(msg.Body))
	}
	return c.JSON(http.StatusOK, "consumed messages from RabbitMQ")
}

func RabbitmqApi() RabbitMQApiInf {
	return &rabbitmqApi{}
}
