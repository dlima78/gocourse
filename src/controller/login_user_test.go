package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/controller/model/request"
	responseModel "github.com/dlima78/gocourse/src/controller/model/response"
	"github.com/dlima78/gocourse/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginUserController_ValidationError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userLoginRequest := request.UserLoginRequest{
		Email:    "wrongemail.com",
		Password: "TEste@123",
	}

	body, _ := json.Marshal(userLoginRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)

	uc := NewUserController(mockService)
	uc.LoginUser(context)

	var errResp rest_err.RestErr
	err := json.Unmarshal(recorder.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, errResp.Code)

	mockService.AssertNotCalled(t, "LoginUserService")
}

func TestLoginUserController_ServerError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userLoginRequest := request.UserLoginRequest{
		Email:    "test@email.com",
		Password: "TEste@123",
	}

	body, _ := json.Marshal(userLoginRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)

	error := rest_err.NewInternalServerError("Error trying to create user")

	mockService.On("LoginUserService", mock.Anything).Return(nil, "", error)

	uc := NewUserController(mockService)
	uc.LoginUser(context)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	mockService.AssertExpectations(t)
}

func TestLoginUserController_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userLoginRequest := request.UserLoginRequest{
		Email:    "test@email.com",
		Password: "TEste@123",
	}

	body, _ := json.Marshal(userLoginRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)

	domainResult := model.NewUserDomain("test@email.com", "TEste@123", "João", 45)
	userID := domainResult.GetID()
	token := "token-jwt-example"

	mockService.On("LoginUserService", mock.Anything).
		Return(domainResult, token, nil)

	uc := NewUserController(mockService)
	uc.LoginUser(context)

	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.EqualValues(t, token, recorder.Header().Values("Authorization")[0], token)

	var resp responseModel.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, userID, resp.ID)
	assert.Equal(t, "test@email.com", resp.Email)
	assert.Equal(t, "João", resp.Name)

	mockService.AssertExpectations(t)
}
