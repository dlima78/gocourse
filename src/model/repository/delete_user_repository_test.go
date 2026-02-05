package repository

import (
	"net/http"
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	t.Run("invalid_id", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		// Usar um ID inválido (não é um ObjectID válido)
		invalidID := "invalid-id-format"

		restErr := repo.DeleteUser(invalidID)

		require.NotNil(t, restErr, "Expected error when using invalid ID format")
		assert.Equal(t, http.StatusBadRequest, restErr.Code, "Expected 400 error code for bad request")

	})

	t.Run("success", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
		createdUser, _ := repo.CreateUser(user)

		restErr := repo.DeleteUser(createdUser.GetID())

		require.Nil(t, restErr, "Expected no error when deleting user")

		result, restErr := repo.FindUserByEmail("test2@mail.com")
		require.NotNil(t, restErr, "Expected error when finding non-existent user by email")
		require.Nil(t, result, "Expected no user to be found for non-existent email")
	})
}
