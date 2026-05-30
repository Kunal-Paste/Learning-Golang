package db

import (
	"context"
	"fmt"
	"note-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {

	// prevent app freze in startup
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoUrl)

	client, err := mongo.Connect(cxt, clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("connect db failed")
	}

	if err := client.Ping(cxt, nil); err != nil {
		return nil, nil, fmt.Errorf("mongo ping failed")
	}

	database := client.Database(cfg.MongoDb)

	return client, database, nil

}

func Disconnect(client *mongo.Client) error {
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Disconnect(cxt)
}
