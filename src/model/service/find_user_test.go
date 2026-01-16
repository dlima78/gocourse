package service

import (
	"testing"

	"github.com/dlima78/gocourse/src/model"
	mock_user_repository "github.com/dlima78/gocourse/src/model/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindUserByEmailService(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	// Configurar mock: FindUserByEmail retorna um usuário existente
	existingUser := model.NewUserDomain("test@example.com", "password123", "João", 30)
	mockRepo.On("FindUserByEmail", "test@example.com").Return(existingUser, nil)

	svc := NewUserDomainService(mockRepo)
	result, err := svc.FindUserByEmailService("test@example.com")

	require.Nil(t, err, "Should not have error")
	require.NotNil(t, result, "Should return user")
	assert.Equal(t, "test@example.com", result.GetEmail())
	assert.Equal(t, "João", result.GetName())
}
func TestFindUserByIDService(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	// Configurar mock: FindUserByID retorna um usuário existente
	userID := "123456"
	existingUser := model.NewUserDomain("test@example.com", "password123", "João", 30)
	mockRepo.On("FindUserByID", userID).Return(existingUser, nil)

	svc := NewUserDomainService(mockRepo)
	result, err := svc.FindUserByIDService(userID)

	require.Nil(t, err, "Should not have error")
	require.NotNil(t, result, "Should return user")
	assert.Equal(t, "test@example.com", result.GetEmail())
	assert.Equal(t, "João", result.GetName())
}
