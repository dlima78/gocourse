package service

import (
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/model"
	mock_user_repository "github.com/dlima78/gocourse/src/model/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserService(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	userID := "123456"
	existingUser := model.NewUserDomain("test@example.com", "password123", "João", 30)
	mockRepo.On("FindUserByID", userID).Return(existingUser, nil)

	mockRepo.On("UpdateUser", userID, mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "newEmail@example.com" &&
			u.GetName() == "João Silva" &&
			u.GetAge() == 31
	})).Return(nil)

	service := NewUserDomainService(mockRepo)

	// Usuário com dados atualizados
	updatedUser := model.NewUserDomain("newEmail@example.com", "password123", "João Silva", 31)
	err := service.UpdateUserService(userID, updatedUser)

	// Assertions
	require.Nil(t, err)
	// Debug: mostrar chamadas registradas no mock
	t.Logf("mock calls: %#v", mockRepo.Calls)
	mockRepo.AssertCalled(t, "FindUserByID", userID)
	mockRepo.AssertNumberOfCalls(t, "UpdateUser", 1)
}

func TestUpdateUserService_UserNotFound(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	userId := "nonexistent"
	mockRepo.On("FindUserByID", userId).Return(nil, rest_err.NewNotFoundError("user not found"))
	service := NewUserDomainService(mockRepo)
	err := service.UpdateUserService(userId, model.NewUserDomain("newEmail@example.com", "password123", "João Silva", 31))
	require.NotNil(t, err)
}

func TestUpdateUserService_RepositoryError(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository) // ou new(MockUserRepository) conforme seu padrão

	userID := "123456"
	existingUser := model.NewUserDomain("old@example.com", "oldpass", "João", 30)
	mockRepo.On("FindUserByID", userID).Return(existingUser, nil)

	mockErr := rest_err.NewInternalServerError("Database update failed")
	mockRepo.On("UpdateUser", userID, mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "new@example.com" && u.GetName() == "João Silva" && u.GetAge() == 31
	})).Return(mockErr)

	service := NewUserDomainService(mockRepo)
	updatedUser := model.NewUserDomain("new@example.com", "password123", "João Silva", 31)
	err := service.UpdateUserService(userID, updatedUser)

	require.NotNil(t, err)
	assert.Equal(t, "Database update failed", err.Message)

	mockRepo.AssertCalled(t, "FindUserByID", userID)
	mockRepo.AssertCalled(t, "UpdateUser", userID, mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "new@example.com"
	}))
}
