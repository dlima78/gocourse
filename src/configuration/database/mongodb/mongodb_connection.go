package mongodb

import (
	"context"
	"time"

	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitConnextion() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.Info("CONNECTED ON MONGODB")
}
