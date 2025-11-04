package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_CRUD_DB = "MONGODB_CRUD_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_CRUD_DB)

	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
