package db

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/v2/x/mongo/driver/mongocrypt/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {
	connectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clientOpt := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(connectCtx, clientOpt)
	if err != nil {
		return nil, fmt.Errorf("mongo connection failed %w", err)
	}

	database := client.Database(cfg.MongoDBname)

	return &Mongo{
		Client: client,
		DB:     database,
	}, nil

}
