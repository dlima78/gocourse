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

// MockUserRepository é um mock do repositório

// Testes do CreateUserService
func TestCreateUserService_Success(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	// Configurar mock: FindUserByEmail retorna nil (email não existe)
	mockRepo.On("FindUserByEmail", "test@example.com").Return(nil, nil)

	// Configurar mock: CreateUser retorna o usuário criado
	newUser := model.NewUserDomain("test@example.com", "password123", "João", 30)
	mockRepo.On("CreateUser", mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "test@example.com"
	})).Return(newUser, nil)

	// Criar service com mock
	service := NewUserDomainService(mockRepo)

	// Executar o teste
	user := model.NewUserDomain("test@example.com", "password123", "João", 30)
	result, err := service.CreateUserService(user)

	// Assertions
	require.Nil(t, err, "Should not have error")
	require.NotNil(t, result, "Should return user")
	assert.Equal(t, "test@example.com", result.GetEmail())
	assert.Equal(t, "João", result.GetName())

	// Verificar que os mocks foram chamados
	mockRepo.AssertCalled(t, "FindUserByEmail", "test@example.com")
	mockRepo.AssertCalled(t, "CreateUser", mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "test@example.com"
	}))
	mockRepo.AssertNumberOfCalls(t, "FindUserByEmail", 1)
	mockRepo.AssertNumberOfCalls(t, "CreateUser", 1)
}

func TestCreateUserService_DuplicateEmail(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	// Configurar mock: FindUserByEmail retorna um usuário existente
	existingUser := model.NewUserDomain("test@example.com", "existingpass", "Maria", 28)
	mockRepo.On("FindUserByEmail", "test@example.com").Return(existingUser, nil)

	// Criar service com mock
	service := NewUserDomainService(mockRepo)

	// Executar o teste
	user := model.NewUserDomain("test@example.com", "password123", "João", 30)
	result, err := service.CreateUserService(user)

	// Assertions
	require.NotNil(t, err, "Should have error for duplicate email")
	require.Nil(t, result, "Should not return user")
	assert.Equal(t, "Email already exists", err.Message)

	// Verificar que CreateUser NÃO foi chamado
	mockRepo.AssertNotCalled(t, "CreateUser")
	mockRepo.AssertNumberOfCalls(t, "FindUserByEmail", 1)
}

func TestCreateUserService_RepositoryError(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	// Configurar mock: FindUserByEmail retorna nil (email não existe)
	mockRepo.On("FindUserByEmail", "test@example.com").Return(nil, nil)

	// Configurar mock: CreateUser retorna erro
	mockErr := rest_err.NewInternalServerError("Database connection failed")
	mockRepo.On("CreateUser", mock.Anything).Return(nil, mockErr)

	// Criar service com mock
	service := NewUserDomainService(mockRepo)

	// Executar o teste
	user := model.NewUserDomain("test@example.com", "password123", "João", 30)
	result, err := service.CreateUserService(user)

	// Assertions
	require.NotNil(t, err, "Should have error from repository")
	require.Nil(t, result, "Should not return user")
	assert.Equal(t, "Database connection failed", err.Message)

	mockRepo.AssertCalled(t, "CreateUser", mock.Anything)
}

func TestCreateUserService_PasswordEncrypted(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	mockRepo.On("FindUserByEmail", "test@example.com").Return(nil, nil)

	newUser := model.NewUserDomain("test@example.com", "encrypted_password", "João", 30)
	mockRepo.On("CreateUser", mock.Anything).Return(newUser, nil)

	service := NewUserDomainService(mockRepo)

	user := model.NewUserDomain("test@example.com", "password123", "João", 30)
	result, err := service.CreateUserService(user)

	// Verificar que a senha foi criptografada (não é a mesma que foi passada)
	require.Nil(t, err)
	require.NotNil(t, result)
	// A senha retornada é a criptografada
	assert.NotEqual(t, "password123", result.GetPassword())

	// Verificar que CreateUser foi chamado com usuário que tem senha criptografada
	mockRepo.AssertCalled(t, "CreateUser", mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetPassword() != "password123" // A senha deve estar criptografada
	}))
}
