package repository

import (
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_TestUpdateUser(t *testing.T) {
	setup := SetupMongoDB(t, "user_database_test")
	defer setup.Cleanup()

	originalValue := os.Getenv("MONGODB_USER_DB")
	os.Setenv("MONGODB_USER_DB", "users")
	defer os.Setenv("MONGODB_USER_DB", originalValue)

	repo := NewUserRepository(setup.Database)

	user := model.NewUserDomain("test@mail.com", "test", "Eduardo", 46)
	createdUser, restErr := repo.CreateUser(user)
	require.Nil(t, restErr, "Expected no error when creating user for find test")

	dataToUpdate := model.NewUserDomain("newemail@mail.com", "newpassword", "Eduardo Lima", 47)

	restErr = repo.UpdateUser(createdUser.GetID(), dataToUpdate)

	require.Nil(t, restErr, "Expected no error when updating user")

	updatedUser, restErr := repo.FindUserByID(createdUser.GetID())

	require.Nil(t, restErr, "Expected no error when finding updated user")
	require.Equal(t, "newemail@mail.com", updatedUser.GetEmail(), "Expected email to match")
	require.Equal(t, "newpassword", updatedUser.GetPassword(), "Expected password to match")
	require.Equal(t, "Eduardo Lima", updatedUser.GetName(), "Expected name to match")
	require.Equal(t, int8(47), updatedUser.GetAge(), "Expected age to match")

}

func TestUserRepository_TestUpdateUser_InvalidID(t *testing.T) {
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

}
