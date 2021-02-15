package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoDBConnection get connection of mongodb
func GetMongoDBConnection() (*mongo.Client, error) {
	// Run the following command first in order to be able to connect to mongodb locally!
	// docker run --name mongo-db -p 27017:27017 -d mongo:latest
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

func GetMongoDBCollection(DBName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDBConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DBName).Collection(CollectionName)

	return collection, nil
}
