package repository

import (
	"net/http"
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	t.Run("user_not_found", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		result, restErr := repo.FindUserByEmail("test@mail.com")
		require.NotNil(t, restErr, "Expected error when finding non-existent user by email")
		require.Nil(t, result, "Expected no user to be found for non-existent email")
	})
	t.Run("success", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
		_, restErr := repo.CreateUser(user)

		require.Nil(t, restErr, "Expected no error when creating user for find test")

		result, restErr := repo.FindUserByEmail("test@mail.com")

		require.Nil(t, restErr, "Expected no error when finding user by email")
		require.NotNil(t, result, "Expected to find a user")
		require.Equal(t, "test@mail.com", result.GetEmail(), "Expected email to match")

		result2, restErr2 := repo.FindUserByEmail("test2@mail.com")
		require.NotNil(t, restErr2, "Expected error when finding non-existent user by email")
		require.Nil(t, result2, "Expected no user to be found for non-existent email")
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	t.Run("invalid_id", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		invalidID := "invalid-id-format"

		result, restErr := repo.FindUserByID(invalidID)

		require.NotNil(t, restErr, "Expected error when using invalid ID format")
		require.Nil(t, result, "Expected no user to be found")
		require.Equal(t, 400, restErr.Code, "Expected 400 error code for bad request")
	})

	t.Run("user_not_found_with_this_id", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()
		id := bson.NewObjectID().Hex()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		result, restErr := repo.FindUserByID(id)

		require.NotNil(t, restErr, "User not found with this ID")
		require.Nil(t, result, "Expected no user to be found")
		require.Equal(t, http.StatusNotFound, restErr.Code,
			"Expected 404 error code for bad request")
	})

	t.Run("success", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
		createdUser, restErr := repo.CreateUser(user)

		require.Nil(t, restErr, "Expected no error when creating user for find test")

		result, restErr := repo.FindUserByID(createdUser.GetID())

		require.Nil(t, restErr, "Expected no error when finding user by id")
		require.NotNil(t, result, "Expected to find a user")
		require.EqualValues(t, "test@mail.com", result.GetEmail(), "Expected email to match")
		require.EqualValues(t, "test", result.GetPassword(), "Expected password to match")
		require.EqualValues(t, "Eduardo", result.GetPassword(), "Expected password to match")

	})

}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	t.Run("user_or_password_invalid", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		// Setar a variável de ambiente para o nome da coleção
		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		result, restErr := repo.FindUserByEmailAndPassword("test@mail.com", "Tes@125esd")

		require.NotNil(t, restErr, "Expected error when finding user by email and password")
		require.Nil(t, result, "Expected to find a user")
		require.Equal(t, 401, restErr.Code, "Expected 401 error code for unauthorized")
	})

	t.Run("succes", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		// Setar a variável de ambiente para o nome da coleção
		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
		_, restErr := repo.CreateUser(user)

		require.Nil(t, restErr, "Expected no error when creating user for find test")

		result, restErr := repo.FindUserByEmailAndPassword("test@mail.com", "test")

		require.Nil(t, restErr, "Expected no error when finding user by email and password")
		require.NotNil(t, result, "Expected to find a user")
		require.Equal(t, "test@mail.com", result.GetEmail(), "Expected email to match")
		require.Equal(t, "test", result.GetPassword(), "Expected password to match")
		require.Equal(t, "Eduardo", result.GetName(), "Expected name to match")
		require.Equal(t, int8(46), result.GetAge(), "Expected age to match")
	})

}
