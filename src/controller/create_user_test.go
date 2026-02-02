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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

func TestCreateUserController_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userRequest := request.UserRequest{
		Email:    "test@mail.com",
		Password: "TEste@123",
		Name:     "João",
		Age:      45,
	}

	body, _ := json.Marshal(userRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)

	var captured model.UserDomainInterface
	createdUser := model.NewUserDomain("test@mail.com", "TEste@123", "João", 45)
	userID := createdUser.GetID()
	mockService.On("CreateUserService", mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "test@mail.com" && u.GetName() == "João"
	})).
		Return(createdUser, nil).
		Run(func(args mock.Arguments) {
			captured = args.Get(0).(model.UserDomainInterface)
		})

	uc := NewUserController(mockService)
	uc.CreateUser(context)

	assert.Equal(t, "test@mail.com", captured.GetEmail())
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var resp responseModel.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, userID, resp.ID)
	assert.Equal(t, "test@mail.com", resp.Email)
	assert.Equal(t, "João", resp.Name)

	mockService.AssertExpectations(t)
}

func TestCreateUserController_ValidationError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userRequest := request.UserRequest{
		Email:    "wrongemail.com",
		Password: "TEste@123",
		Name:     "João",
		Age:      45,
	}

	body, _ := json.Marshal(userRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)

	uc := NewUserController(mockService)
	uc.CreateUser(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var errResp rest_err.RestErr
	err := json.Unmarshal(recorder.Body.Bytes(), &errResp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, errResp.Code)
	// validar que o serviço NÃO foi chamado
	mockService.AssertNotCalled(t, "CreateUserService")
}

func TestCreateUserController_DuplicateEmail(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userRequest := request.UserRequest{
		Email:    "test@mail.com",
		Password: "TEste@123",
		Name:     "João",
		Age:      45,
	}

	body, _ := json.Marshal(userRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))
	mock_user_controller.MakeRequest(
		context,
		gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)
	errResp := rest_err.NewBadRequestError("Email already exists")
	mockService.On("CreateUserService", mock.Anything).Return(nil, errResp)

	uc := NewUserController(mockService)
	uc.CreateUser(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)

	var resp rest_err.RestErr
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Email already exists", resp.Message)

	mockService.AssertExpectations(t)
}

func TestCreateUserController_ServerError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	userRequest := request.UserRequest{
		Email:    "test@mail.com",
		Password: "TEste@123",
		Name:     "João",
		Age:      45,
	}

	body, _ := json.Marshal(userRequest)
	mock_user_controller.MakeRequest(
		context, gin.Params{},
		url.Values{},
		"POST",
		io.NopCloser(bytes.NewBufferString(string(body))))

	mockService := new(mock_user_controller.MockUserService)

	error := rest_err.NewInternalServerError("Error trying to create user")

	mockService.On("CreateUserService", mock.Anything).Return(nil, error)

	uc := NewUserController(mockService)
	uc.CreateUser(context)

	assert.Equal(t, 500, recorder.Code)

}
