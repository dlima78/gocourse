package repository

import (
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_CreateUser(t *testing.T) {
	t.Run("duplicated_email", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		user1 := model.NewUserDomain("duplicate@mail.com", "password123", "João", 30)
		result1, err1 := repo.CreateUser(user1)
		require.Nil(t, err1, "First user should be created without error")
		require.NotNil(t, result1, "First user should not be nil")

		user2 := model.NewUserDomain("duplicate@mail.com", "password456", "Maria", 28)
		result2, err2 := repo.CreateUser(user2)

		assert.NotNil(t, err2, "Should have error when creating user with duplicate email")
		assert.Nil(t, result2, "Result should be nil when there is an error")
	})

	t.Run("success", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		// Setar a variável de ambiente para o nome da coleção
		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		// Criar repositório com a conexão real
		repo := NewUserRepository(setup.Database)

		// Executar o teste
		user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
		result, restErr := repo.CreateUser(user)

		// Assertions
		require.Nil(t, restErr, "Expected no error when creating user")
		require.NotNil(t, result, "Expected user to be created")
		assert.Equal(t, "test@mail.com", result.GetEmail())
		assert.Equal(t, "Eduardo", result.GetName())
		assert.Equal(t, int8(46), result.GetAge())
	})

}
