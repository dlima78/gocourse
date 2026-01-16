package service

import (
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	mock_user_repository "github.com/dlima78/gocourse/src/model/service/mocks"
)

func TestDeleteUserService_Success(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	mockRepo.On("DeleteUser", "user123").Return(nil)

	service := NewUserDomainService(mockRepo)

	err := service.DeleteUserService("user123")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockRepo.AssertCalled(t, "DeleteUser", "user123")
	mockRepo.AssertNumberOfCalls(t, "DeleteUser", 1)

}

func TestDeleteUserService_Error(t *testing.T) {
	mockRepo := new(mock_user_repository.MockUserRepository)

	mockErr := &rest_err.RestErr{
		Message: "User not found",
		Code:    404,
	}

	mockRepo.On("DeleteUser", "user123").Return(mockErr)

	service := NewUserDomainService(mockRepo)

	err := service.DeleteUserService("user123")

	if err == nil {
		t.Errorf("Expected error, got nil")
	} else if err.Message != "User not found" {
		t.Errorf("Expected error message 'User not found', got %v", err.Message)
	}

	mockRepo.AssertCalled(t, "DeleteUser", "user123")
	mockRepo.AssertNumberOfCalls(t, "DeleteUser", 1)
}
