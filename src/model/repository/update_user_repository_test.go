package repository

import (
	"net/http"
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_TestUpdateUser(t *testing.T) {
	t.Run("invalid_id", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)
		user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)

		_, restErr := repo.CreateUser(user)
		require.Nil(t, restErr, "Expected no error when creating user for find test")

		dataToUpdate := model.NewUserDomain("newemail@mail.com", "newpassword", "Eduardo Lima", 47)

		invalidID := "invalid-id-format"
		restErr = repo.UpdateUser(invalidID, dataToUpdate)

		require.NotNil(t, restErr, "Expected error when updating user with invalid ID")
		require.Equal(t, http.StatusBadRequest, restErr.Code, "Expected error code 400 for bad request")
	})

	t.Run("success", func(t *testing.T) {
		setup := SetupMongoDB(t, "user_database_test")
		defer setup.Cleanup()

		originalValue := os.Getenv("MONGODB_USER_DB")
		os.Setenv("MONGODB_USER_DB", "users")
		defer os.Setenv("MONGODB_USER_DB", originalValue)

		repo := NewUserRepository(setup.Database)

		user := model.NewUserDomain("test@mail.com", "TESs@123", "Eduardo", 46)
		createdUser, restErr := repo.CreateUser(user)
		require.Nil(t, restErr, "Expected no error when creating user for find test")

		dataToUpdate := model.NewUserDomain("newemail@mail.com", "NewPassword@123", "Eduardo Lima", 47)

		restErr = repo.UpdateUser(createdUser.GetID(), dataToUpdate)

		require.Nil(t, restErr, "Expected no error when updating user")

		updatedUser, restErr := repo.FindUserByID(createdUser.GetID())

		require.Nil(t, restErr, "Expected no error when finding updated user")
		require.Equal(t, "newemail@mail.com", updatedUser.GetEmail(), "Expected email to match")
		require.Equal(t, "NewPassword@123", updatedUser.GetPassword(), "Expected password to match")
		require.Equal(t, "Eduardo Lima", updatedUser.GetName(), "Expected name to match")
		require.Equal(t, int8(47), updatedUser.GetAge(), "Expected age to match")
	})

}
