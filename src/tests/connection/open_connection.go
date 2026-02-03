package connection

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func OpenConnection() (database *mongo.Database, close func()) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
		return
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})
	if err != nil {
		log.Fatalf("Could not create mongo container: %s", err)
		return
	}

	var client *mongo.Client
	uri := fmt.Sprintf("mongodb://127.0.0.1:%s", resource.GetPort("27017/tcp"))

	if err := pool.Retry(func() error {
		var err error

		client, err = mongo.Connect(options.Client().ApplyURI(uri))
		if err != nil {
			return err
		}

		return client.Ping(context.TODO(), nil)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	dbName := os.Getenv("MONGODB_USER_DB")
	if dbName == "" {
		dbName = "test_db"
	}

	database = client.Database(dbName)

	close = func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Printf("Error disconnecting mongo: %s", err)
		}
		if err := pool.Purge(resource); err != nil {
			log.Printf("Error purging resource: %s", err)
		}
	}

	return
}
