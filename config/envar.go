package config

import (
	"log"
	"os"
	"strings"
)

var RunMode string
var ServerPort string
var MongoDbConnectionStringForWrite string
var MongoDbConnectionStringForRead string
var DatabaseName string

func InitEnvironmentVariables() {
	RunMode = strings.TrimSpace(os.Getenv("RUN_MODE"))
	if RunMode == "" {
		RunMode = "DEVELOP"
	}

	ServerPort = strings.TrimSpace(os.Getenv("SERVER_PORT"))
	if ServerPort == "" {
		ServerPort = "8080"
	}

	MongoDbConnectionStringForWrite = strings.TrimSpace(os.Getenv("MONGODB_CONNECTION_STRING_FOR_WRITE"))
	MongoDbConnectionStringForRead = strings.TrimSpace(os.Getenv("MONGODB_CONNECTION_STRING_FOR_READ"))
	DatabaseName = strings.TrimSpace(os.Getenv("DATABASE_NAME"))

	log.Println("Run Mode: " + RunMode)
	log.Println("Server Port: " + ServerPort)
	log.Println("Database Name: " + DatabaseName)

	if MongoDbConnectionStringForWrite != "" {
		log.Println("Mongo Conn String For Write: " + MongoDbConnectionStringForWrite)
		log.Println("Mongo Conn String For Read: " + MongoDbConnectionStringForRead)
	}

}
