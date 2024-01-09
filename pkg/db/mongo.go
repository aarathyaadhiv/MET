package db

import (
	"context"
	"fmt"

	"github.com/aarathyaadhiv/met/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongo(c config.Config) (*mongo.Client, error) {
	ctx := context.TODO()
	mongoConn := options.Client().ApplyURI(c.DB_URL)
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		return nil, err
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("mongo connection established")
	return mongoClient, nil
}
