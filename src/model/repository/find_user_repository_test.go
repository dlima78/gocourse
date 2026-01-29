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

func TestUserRepository_FindUserByID(t *testing.T) {
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
	createdUser, restErr := repo.CreateUser(user)

	require.Nil(t, restErr, "Expected no error when creating user for find test")

	// Executar o teste de busca usando o ID do usuário criado
	result, restErr := repo.FindUserByID(createdUser.GetID())

	// Assertions
	require.Nil(t, restErr, "Expected no error when finding user by id")
	require.NotNil(t, result, "Expected to find a user")
	require.Equal(t, "test@mail.com", result.GetEmail(), "Expected email to match")

	// result2, restErr2 := repo.FindUserByEmail("test2@mail.com")
	// require.NotNil(t, restErr2, "Expected error when finding non-existent user by email")
	// require.Nil(t, result2, "Expected no user to be found for non-existent email")
}

func TestUserRepository_FindUserByID_InvalidID(t *testing.T) {
	// Setup MongoDB com testcontainers
	setup := SetupMongoDB(t, "user_database_test")
	defer setup.Cleanup()

	// Setar a variável de ambiente para o nome da coleção
	originalValue := os.Getenv("MONGODB_USER_DB")
	os.Setenv("MONGODB_USER_DB", "users")
	defer os.Setenv("MONGODB_USER_DB", originalValue)

	repo := NewUserRepository(setup.Database)

	// Usar um ID inválido (não é um ObjectID válido)
	invalidID := "invalid-id-format"

	result, restErr := repo.FindUserByID(invalidID)

	// Assertions
	require.NotNil(t, restErr, "Expected error when using invalid ID format")
	require.Nil(t, result, "Expected no user to be found")
	require.Equal(t, 400, restErr.Code, "Expected 400 error code for bad request")
}

func TestUserRepository_FindUserByEmailAndPassword_Success(t *testing.T) {
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

}

func TestUserRepository_FindUserByEmailAndPassword_Error(t *testing.T) {
	setup := SetupMongoDB(t, "user_database_test")
	defer setup.Cleanup()

	// Setar a variável de ambiente para o nome da coleção
	originalValue := os.Getenv("MONGODB_USER_DB")
	os.Setenv("MONGODB_USER_DB", "users")
	defer os.Setenv("MONGODB_USER_DB", originalValue)

	repo := NewUserRepository(setup.Database)

	result, restErr := repo.FindUserByEmailAndPassword("test@mail.com", "test")

	require.NotNil(t, restErr, "Expected error when finding user by email and password")
	require.Nil(t, result, "Expected to find a user")
	require.Equal(t, 401, restErr.Code, "Expected 401 error code for unauthorized")

}
