package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetConnect() *mongo.Client {
	option := options.Client().ApplyURI("mongodb://114.55.101.215:27017/test")
	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	return client
}
