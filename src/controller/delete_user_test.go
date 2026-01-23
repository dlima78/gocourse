package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

// MockUserService implementa a interface service.UserDomainService apenas para os testes
// métodos que não são usados nos testes retornam zeros/nil para simplificar

func TestDeleteUserController_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)
	param := gin.Params{gin.Param{Key: "userId", Value: "123"}}

	mock_user_controller.MakeRequest(context, param, url.Values{}, "DELETE", nil)

	mockService := new(mock_user_controller.MockUserService)
	mockService.On("DeleteUserService", "123").Return(nil)

	uc := NewUserController(mockService)
	uc.DeleteUser(context)

	assert.EqualValues(t, http.StatusOK, recorder.Code)
}

func TestDeleteUserController_NotFound(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)
	params := gin.Params{gin.Param{Key: "userId", Value: "123"}}

	mock_user_controller.MakeRequest(context, params, url.Values{}, "DELETE", nil)

	mockService := new(mock_user_controller.MockUserService)
	errResp := rest_err.NewNotFoundError("user not found")
	mockService.On("DeleteUserService", "123").Return(errResp)

	uc := NewUserController(mockService)
	uc.DeleteUser(context)

	assert.Equal(t, http.StatusNotFound, recorder.Code)

	var resp rest_err.RestErr
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, errResp.Message, resp.Message)
	mockService.AssertExpectations(t)
}

func TestDeleteUserController_InternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)
	params := gin.Params{gin.Param{Key: "userId", Value: "123"}}

	mock_user_controller.MakeRequest(context, params, url.Values{}, "DELETE", nil)

	mockService := new(mock_user_controller.MockUserService)
	errResp := rest_err.NewInternalServerError("database error")
	mockService.On("DeleteUserService", "123").Return(errResp)

	uc := NewUserController(mockService)
	uc.DeleteUser(context)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)

	var resp rest_err.RestErr
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, errResp.Message, resp.Message)
	mockService.AssertExpectations(t)
}
