package MongoDA

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoDbConnection get connection of mongodb
func GetMongoDbConnection() (*mongo.Client, error) {

	/*Mongo is running on port 27017 on docker change according to your port*/
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
