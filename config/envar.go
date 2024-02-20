package config

import (
	"log"
	"os"
	"strings"
)

var RunMode string
var ServerPort string
var ConnectMongo string
var MongoDbConnectionStringForWrite string
var MongoDbConnectionStringForRead string
var DatabaseName string

var ConnectRedis string
var RedisServerForWrite string
var RedisServerForRead string
var RedisServerPassword string

var ConnectRabbitMQ string
var RabbitMQUser string
var RabbitMQPassword string
var RabbitMQServer string
var RabbitMQConnectionUrl string

func InitEnvironmentVariables() {
	RunMode = strings.TrimSpace(os.Getenv("RUN_MODE"))
	if RunMode == "" {
		RunMode = "DEVELOP"
	}

	ServerPort = strings.TrimSpace(os.Getenv("SERVER_PORT"))
	if ServerPort == "" {
		ServerPort = "8080"
	}

	ConnectMongo = strings.TrimSpace(os.Getenv("CONNECT_MONGO"))
	MongoDbConnectionStringForWrite = strings.TrimSpace(os.Getenv("MONGODB_CONNECTION_STRING_FOR_WRITE"))
	MongoDbConnectionStringForRead = strings.TrimSpace(os.Getenv("MONGODB_CONNECTION_STRING_FOR_READ"))
	DatabaseName = strings.TrimSpace(os.Getenv("DATABASE_NAME"))

	ConnectRedis = strings.TrimSpace(os.Getenv("CONNECT_REDIS"))
	RedisServerForWrite = strings.TrimSpace(os.Getenv("REDIS_SERVER_FOR_WRITE"))
	RedisServerForRead = strings.TrimSpace(os.Getenv("REDIS_SERVER_FOR_READ"))
	RedisServerPassword = strings.TrimSpace(os.Getenv("REDIS_SERVER_PASSWORD"))

	ConnectRabbitMQ = strings.TrimSpace(os.Getenv("CONNECT_RABBITMQ"))
	RabbitMQUser = strings.TrimSpace(os.Getenv("RABBITMQ_USER"))
	RabbitMQPassword = strings.TrimSpace(os.Getenv("RABBITMQ_PASSWORD"))
	RabbitMQServer = strings.TrimSpace(os.Getenv("RABBITMQ_SERVER"))
	if RabbitMQUser != "" && RabbitMQPassword != "" {
		RabbitMQConnectionUrl = "amqp://" + RabbitMQUser + ":" + RabbitMQPassword + "@" + RabbitMQServer + ":5672"
	}

	log.Println("Run Mode: " + RunMode)
	log.Println("Server Port: " + ServerPort)
	log.Println("Mongo Connect: " + ConnectMongo)
	log.Println("Redis Connect: " + ConnectRedis)
	log.Println("Rabbitmq Connect: " + ConnectRabbitMQ)

	if ConnectMongo == "true" {
		log.Println("Mongo Conn String For Write: " + MongoDbConnectionStringForWrite)
		log.Println("Mongo Conn String For Read: " + MongoDbConnectionStringForRead)
	}

	if ConnectRedis == "true" {
		log.Println("Redis Server For Write: " + RedisServerForWrite)
		log.Println("Mongo Server For Read: " + RedisServerForRead)
		log.Println("Mongo Server Password: " + RedisServerPassword)
	}

	if ConnectRabbitMQ == "true" {
		log.Println("connection string for rabbitmq: " + RabbitMQConnectionUrl)
		log.Println("rabbitMQ server: " + RabbitMQServer)
		log.Println("rabbitMQ user: " + RabbitMQUser)
		log.Println("rabbitMQ password: " + RabbitMQPassword)
	}
}
