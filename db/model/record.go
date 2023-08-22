package model

import (
	"app-container-platform/db/cp_mongodb"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type RecordConfig struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string `bson:"name" json:"name"`
	UID                int    `bson:"uid" json:"uid"`
}

func (db RecordConfig) SaveToMongo() error {
	return cp_mongodb.RecordDBCollectionForWrite.Save(&db)
}

func (db RecordConfig) GetListFromMongo(query map[string]interface{}) ([]RecordConfig, error) {
	dbList := []RecordConfig{}
	results := cp_mongodb.RecordDBCollectionForRead.Find(query)
	existingDatabase := &RecordConfig{}
	for results.Next(existingDatabase) {
		dbList = append(dbList, *existingDatabase)
	}
	return dbList, nil
}

func (db RecordConfig) GetByIdFromMongo(id bson.ObjectId) (RecordConfig, error) {
	query := bson.M{"$and": []bson.M{
		bson.M{"_id": id},
	},
	}
	existingDatabase := &RecordConfig{}
	err := cp_mongodb.RecordDBCollectionForRead.FindOne(query, existingDatabase)
	if err != nil {
		return RecordConfig{}, nil
	}
	return *existingDatabase, nil
}

func (db RecordConfig) GetByUIdFromMongo(uid string) (RecordConfig, error) {
	query := bson.M{"$and": []bson.M{
		bson.M{"uid": uid},
	},
	}
	existingDatabase := &RecordConfig{}
	err := cp_mongodb.RecordDBCollectionForRead.FindOne(query, existingDatabase)
	if err != nil {
		return RecordConfig{}, nil
	}
	return *existingDatabase, nil
}

func (db RecordConfig) DeleteByIdFromMongo(id bson.ObjectId) error {
	query := bson.M{"$and": []bson.M{
		bson.M{"_id": id},
	},
	}
	err := cp_mongodb.RecordDBCollectionForWrite.DeleteOne(query)
	if err != nil {
		return err
	}
	return nil
}
