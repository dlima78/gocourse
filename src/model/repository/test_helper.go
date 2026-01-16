package repository

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// MongoTestSetup contém a configuração do MongoDB para testes
type MongoTestSetup struct {
	Client    *mongo.Client
	Database  *mongo.Database
	Container *mongodb.MongoDBContainer
	Context   context.Context
}

// SetupMongoDB cria um container MongoDB e retorna a conexão pronta para usar
func SetupMongoDB(t *testing.T, databaseName string) *MongoTestSetup {
	ctx := context.Background()

	// Iniciar container MongoDB
	mongodbContainer, err := mongodb.Run(ctx, "mongo:6")
	if err != nil {
		t.Fatalf("Failed to start MongoDB container: %v", err)
	}

	// Obter a string de conexão
	mongoURI, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		mongodbContainer.Terminate(ctx)
		t.Fatalf("Failed to get connection string: %v", err)
	}

	// Conectar ao MongoDB
	opts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(opts)
	if err != nil {
		mongodbContainer.Terminate(ctx)
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Testar a conexão
	if err := client.Ping(ctx, nil); err != nil {
		client.Disconnect(ctx)
		mongodbContainer.Terminate(ctx)
		t.Fatalf("Failed to ping MongoDB: %v", err)
	}

	database := client.Database(databaseName)

	return &MongoTestSetup{
		Client:    client,
		Database:  database,
		Container: mongodbContainer,
		Context:   ctx,
	}
}

// Cleanup encerra a conexão e termina o container
func (m *MongoTestSetup) Cleanup() error {
	if err := m.Client.Disconnect(m.Context); err != nil {
		return err
	}
	return m.Container.Terminate(m.Context)
}
