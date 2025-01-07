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
var RedisConnectionType string // MASTER_SLAVE / SENTINEL / CLUSTER
var RedisSentinelServer string
var RedisSentinelMasterName string
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
	RedisConnectionType = strings.TrimSpace(os.Getenv("REDIS_CONNECTION_TYPE")) // MASTER_SLAVE / SENTINEL / CLUSTER
	RedisServerForWrite = strings.TrimSpace(os.Getenv("REDIS_SERVER_FOR_WRITE"))
	RedisServerForRead = strings.TrimSpace(os.Getenv("REDIS_SERVER_FOR_READ"))
	RedisServerPassword = strings.TrimSpace(os.Getenv("REDIS_SERVER_PASSWORD"))

	RedisSentinelServer = strings.TrimSpace(os.Getenv("REDIS_SENTINEL_SERVER"))
	RedisSentinelMasterName = strings.TrimSpace(os.Getenv("REDIS_SENTINEL_MASTER_NAME"))

	ConnectRabbitMQ = strings.TrimSpace(os.Getenv("CONNECT_RABBITMQ"))
	RabbitMQUser = strings.TrimSpace(os.Getenv("RABBITMQ_USER"))
	RabbitMQPassword = strings.TrimSpace(os.Getenv("RABBITMQ_PASSWORD"))
	RabbitMQServer = strings.TrimSpace(os.Getenv("RABBITMQ_SERVER"))
	if RabbitMQUser != "" && RabbitMQPassword != "" {
		RabbitMQConnectionUrl = "amqp://" + RabbitMQUser + ":" + RabbitMQPassword + "@" + RabbitMQServer + ":5672/"
	}

	if RedisSentinelMasterName == "" {
		RedisSentinelMasterName = "mymaster"
	}

	if ConnectMongo == "" {
		ConnectMongo = "false"
	}
	if ConnectRedis == "" {
		ConnectRedis = "false"
	}
	if ConnectRabbitMQ == "" {
		ConnectRabbitMQ = "false"
	}

	log.Println("Run Mode: " + RunMode)
	log.Println("Server Port: " + ServerPort)
	log.Println("Mongo Connect: " + ConnectMongo)
	log.Println("Redis Connect: " + ConnectRedis)
	log.Println("Rabbitmq Connect: " + ConnectRabbitMQ)

	if ConnectMongo == "true" {
		if MongoDbConnectionStringForWrite == "" || MongoDbConnectionStringForRead == "" || DatabaseName == "" {
			log.Println("Unable to connect to Mongo! Missing required values")
			os.Exit(0)
		} else {
			log.Println("Mongo Connection String For Write: " + MongoDbConnectionStringForWrite)
			log.Println("Mongo Connection String For Read: " + MongoDbConnectionStringForRead)
			log.Println("Database Name: " + DatabaseName)
		}
	}

	if ConnectRedis == "true" {
		if RedisConnectionType == "" {
			log.Println("Unable to connect to Redis! Missing connect Type")
			os.Exit(0)
		}

		log.Println("Redis Connect Type: " + RedisConnectionType)
		if RedisConnectionType == "SENTINEL" {
			if RedisSentinelServer == "" || RedisSentinelMasterName == "" {
				log.Println("Unable to connect to Redis! Missing required values")
				os.Exit(0)
			}
			log.Println("Redis Sentinel Server: " + RedisSentinelServer)
			log.Println("Redis Server Master Name: " + RedisSentinelMasterName)
		} else {
			if RedisServerForWrite == "" || RedisServerForRead == "" {
				log.Println("Unable to connect to Redis! Missing required values")
				os.Exit(0)
			}
			log.Println("Redis Server For Write: " + RedisServerForWrite)
			log.Println("Redis Server For Read: " + RedisServerForRead)
		}

		if RedisServerPassword != "" {
			log.Println("Redis Server Password: " + RedisServerPassword)
		}
	}

	if ConnectRabbitMQ == "true" {
		log.Println("connection string for rabbitmq: " + RabbitMQConnectionUrl)
		log.Println("rabbitMQ server: " + RabbitMQServer)
		log.Println("rabbitMQ user: " + RabbitMQUser)
		log.Println("rabbitMQ password: " + RabbitMQPassword)
	}
}
