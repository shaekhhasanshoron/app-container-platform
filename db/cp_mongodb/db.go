package cp_mongodb

import (
	"app-container-platform/config"
	"github.com/go-bongo/bongo"
	"log"
)

var writeConnectDB *bongo.Connection
var readConnectDB *bongo.Connection
var RecordDBCollectionForWrite *bongo.Collection
var RecordDBCollectionForRead *bongo.Collection

func InitMongoDbWriteConnection() error {
	connection, err := createConnectionDB(config.MongoDbConnectionStringForWrite, config.DatabaseName)
	if err != nil {
		log.Println("[ERROR] While initializing MongoDB Write connection:", err.Error())
		return err
	}
	writeConnectDB = connection

	log.Println("[INFO] MongoDB Write connection initialized")
	return nil
}

func InitMongoDbReadConnection() error {
	connection, err := createConnectionDB(config.MongoDbConnectionStringForRead, config.DatabaseName)
	if err != nil {
		log.Println("[ERROR] While initializing MongoDB Write connection:", err.Error())
		return err
	}
	readConnectDB = connection

	log.Println("[INFO] MongoDB Read connection initialized")
	return nil
}

func InitDBCollections() {
	RecordDBCollectionForWrite = writeConnectDB.Collection("records")
	RecordDBCollectionForRead = readConnectDB.Collection("records")
}

func createConnectionDB(connectionString string, dbName string) (*bongo.Connection, error) {
	config := &bongo.Config{
		ConnectionString: connectionString,
		Database:         dbName,
	}
	connection, err := bongo.Connect(config)
	return connection, err
}
