package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

// MockUserService implementa a interface service.UserDomainService apenas para os testes
// métodos que não são usados nos testes retornam zeros/nil para simplificar
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUserService(u model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(u)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}
func (m *MockUserService) UpdateUserService(id string, u model.UserDomainInterface) *rest_err.RestErr {
	args := m.Called(id, u)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*rest_err.RestErr)
}
func (m *MockUserService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(email)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}
func (m *MockUserService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	args := m.Called(id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*rest_err.RestErr)
	}
	if args.Get(1) == nil {
		return args.Get(0).(model.UserDomainInterface), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.Get(1).(*rest_err.RestErr)
}
func (m *MockUserService) DeleteUserService(id string) *rest_err.RestErr {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*rest_err.RestErr)
}
func (m *MockUserService) LoginUserService(u model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	args := m.Called(u)
	if args.Get(0) == nil {
		if args.Get(2) == nil {
			return nil, "", nil
		}
		return nil, "", args.Get(2).(*rest_err.RestErr)
	}
	if args.Get(2) == nil {
		return args.Get(0).(model.UserDomainInterface), args.String(1), nil
	}
	return args.Get(0).(model.UserDomainInterface), args.String(1), args.Get(2).(*rest_err.RestErr)
}

func TestDeleteUserController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("DELETE", "/users/123", nil)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "userId", Value: "123"}}

	mockService := new(MockUserService)
	mockService.On("DeleteUserService", "123").Return(nil)

	uc := NewUserController(mockService)
	uc.DeleteUser(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
	mockService.AssertExpectations(t)
}

func TestDeleteUserController_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("DELETE", "/users/123", nil)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "userId", Value: "123"}}

	mockService := new(MockUserService)
	errResp := rest_err.NewNotFoundError("user not found")
	mockService.On("DeleteUserService", "123").Return(errResp)

	uc := NewUserController(mockService)
	uc.DeleteUser(c)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var resp rest_err.RestErr
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, errResp.Message, resp.Message)
	mockService.AssertExpectations(t)
}
