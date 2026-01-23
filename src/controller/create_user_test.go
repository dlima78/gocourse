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
	responseModel "github.com/dlima78/gocourse/src/controller/model/response"
	"github.com/dlima78/gocourse/src/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
)

func TestCreateUserController_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	body := `{"email":"test@example.com","password":"pass!23","name":"João","age":30}`
	mock_user_controller.MakeRequest(context, gin.Params{}, url.Values{}, "POST", io.NopCloser(bytes.NewBufferString(body)))

	mockService := new(mock_user_controller.MockUserService)

	var captured model.UserDomainInterface
	createdUser := model.NewUserDomain("test@example.com", "pass!23", "João", 30)
	createdUser.SetID("123")
	mockService.On("CreateUserService", mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetEmail() == "test@example.com" && u.GetName() == "João"
	})).
		Return(createdUser, nil).
		Run(func(args mock.Arguments) {
			captured = args.Get(0).(model.UserDomainInterface)
		})

	uc := NewUserController(mockService)
	uc.CreateUser(context)

	assert.Equal(t, "test@example.com", captured.GetEmail())
	assert.EqualValues(t, http.StatusOK, recorder.Code)

	var resp responseModel.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "123", resp.ID)
	assert.Equal(t, "test@example.com", resp.Email)
	assert.Equal(t, "João", resp.Name)

	mockService.AssertExpectations(t)
}

func TestCreateUserController_ValidationError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	// corpo inválido: falta o campo email
	body := `{"password":"pass!23","name":"João","age":30}`
	mock_user_controller.MakeRequest(context, gin.Params{}, url.Values{}, "POST", io.NopCloser(bytes.NewBufferString(body)))

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

	body := `{"email":"duplicate@example.com","password":"pass!23","name":"João","age":30}`
	mock_user_controller.MakeRequest(context, gin.Params{}, url.Values{}, "POST", io.NopCloser(bytes.NewBufferString(body)))

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
