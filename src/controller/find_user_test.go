package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	responseModel "github.com/dlima78/gocourse/src/controller/model/response"
	"github.com/dlima78/gocourse/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindUserByIDController_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	id := "507f1f77bcf86cd799439011" // 24-char hex id

	mockService := new(mock_user_controller.MockUserService)

	mock_user_controller.MakeRequest(context, gin.Params{gin.Param{Key: "userId", Value: id}}, nil, "GET", nil)
	user := model.NewUserDomain("test@example.com", "pass!23", "João", 30)
	user.SetID(id)

	mockService.On("FindUserByIDService", id).Return(user, nil)

	controller := NewUserController(mockService)
	controller.FindUserByID(context)

	mockService.AssertExpectations(t)

	var resp responseModel.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", resp.Email)
	assert.Equal(t, "João", resp.Name)
	assert.Equal(t, 200, recorder.Code)

}

func TestFindUserByIDController_InvalidObjectID(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := mock_user_controller.GetTestingGinContext(recorder)

	params := gin.Params{gin.Param{Key: "userId", Value: "123"}}
	mock_user_controller.MakeRequest(ctx, params, nil, "GET", nil)

	mockService := new(mock_user_controller.MockUserService)
	uc := NewUserController(mockService)
	uc.FindUserByID(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockService.AssertNotCalled(t, "FindUserByIDService")
}

func TestFindUserByIDController_NotFound(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := mock_user_controller.GetTestingGinContext(recorder)

	id := "507f1f77bcf86cd799439011" // 24-char hex válido
	params := gin.Params{gin.Param{Key: "userId", Value: id}}
	mock_user_controller.MakeRequest(ctx, params, nil, "GET", nil)

	mockService := new(mock_user_controller.MockUserService)
	errResp := rest_err.NewNotFoundError("user not found")
	mockService.On("FindUserByIDService", id).Return(nil, errResp)

	uc := NewUserController(mockService)
	uc.FindUserByID(ctx)

	assert.Equal(t, http.StatusNotFound, recorder.Code)

	var resp rest_err.RestErr
	_ = json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.Equal(t, errResp.Message, resp.Message)

	mockService.AssertExpectations(t)
}

func TestFindUserByEmailController_Success(t *testing.T) {
	email := "est@example.com"
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)
	mockService := new(mock_user_controller.MockUserService)

	mock_user_controller.MakeRequest(context, gin.Params{gin.Param{Key: "userEmail", Value: email}}, nil, "GET", nil)

	user := model.NewUserDomain("test@example.com", "pass!23", "João", 30)
	mockService.On("FindUserByEmailService", email).Return(user, nil)

	controller := NewUserController(mockService)
	controller.FindUserByEmail(context)

	mockService.AssertExpectations(t)

	var resp responseModel.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", resp.Email)
	assert.Equal(t, "João", resp.Name)
	assert.Equal(t, 200, recorder.Code)
}

func TestFindUserByEmailController_InvalidEmail(t *testing.T) {
	email := "invalidemail.com"
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)
	mockService := new(mock_user_controller.MockUserService)

	params := gin.Params{gin.Param{Key: "userEmail", Value: email}}
	mock_user_controller.MakeRequest(context, params, nil, "GET", nil)

	controller := NewUserController(mockService)
	controller.FindUserByEmail(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockService.AssertNotCalled(t, "FindUserByEmailService")

}
