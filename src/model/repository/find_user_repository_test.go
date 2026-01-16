package repository

import (
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	// Setup MongoDB com testcontainers
	setup := SetupMongoDB(t, "user_database_test")
	defer setup.Cleanup()

	// Setar a variável de ambiente para o nome da coleção
	originalValue := os.Getenv("MONGODB_USER_DB")
	os.Setenv("MONGODB_USER_DB", "users")
	defer os.Setenv("MONGODB_USER_DB", originalValue)
	// Criar repositório com a conexão real
	repo := NewUserRepository(setup.Database)

	user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
	_, restErr := repo.CreateUser(user)

	require.Nil(t, restErr, "Expected no error when creating user for find test")

	// Executar o teste de busca
	result, restErr := repo.FindUserByEmail("test@mail.com")

	// Assertions
	require.Nil(t, restErr, "Expected no error when finding user by email")
	require.NotNil(t, result, "Expected to find a user")
	require.Equal(t, "test@mail.com", result.GetEmail(), "Expected email to match")

	result2, restErr2 := repo.FindUserByEmail("test2@mail.com")
	require.NotNil(t, restErr2, "Expected error when finding non-existent user by email")
	require.Nil(t, result2, "Expected no user to be found for non-existent email")
}
