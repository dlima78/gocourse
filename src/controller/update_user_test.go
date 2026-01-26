package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUserController_Success(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	id := "507f1f77bcf86cd799439011" // 24-char hex id
	body := `{"name":"Novo Nome","age":31}`

	mockService := new(mock_user_controller.MockUserService)
	params := gin.Params{gin.Param{Key: "userId", Value: id}}
	mock_user_controller.MakeRequest(
		context,
		params,
		nil,
		"PUT",
		io.NopCloser(strings.NewReader(body)))

	mockService.On("UpdateUserService", id, mock.MatchedBy(func(u model.UserDomainInterface) bool {
		return u.GetName() == "Novo Nome" && u.GetAge() == 31
	})).Return(nil)

	controller := NewUserController(mockService)
	controller.UpdateUser(context)

	mockService.AssertExpectations(t)

	assert.Equal(t, http.StatusOK, recorder.Code)

}

func TestUpdateUserController_ValidationError(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	id := "507f1f77bcf86cd799439011"
	body := `{"name":"a"}`
	mockService := new(mock_user_controller.MockUserService)

	params := gin.Params{gin.Param{Key: "userId", Value: id}}
	mock_user_controller.MakeRequest(
		context,
		params,
		nil,
		"PUT",
		io.NopCloser(strings.NewReader(body)))

	controller := NewUserController(mockService)
	controller.UpdateUser(context)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockService.AssertNotCalled(t, "UpdateUserService")

}

func TestUpdateUserController_InvalidObjectID(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := mock_user_controller.GetTestingGinContext(recorder)
	id := "123"

	params := gin.Params{gin.Param{Key: "userId", Value: id}}
	mock_user_controller.MakeRequest(
		ctx,
		params,
		nil,
		"PUT",
		io.NopCloser(strings.NewReader(`{"name":"X"}`)))

	mockService := new(mock_user_controller.MockUserService)
	controller := NewUserController(mockService)
	controller.UpdateUser(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	mockService.AssertNotCalled(t, "UpdateUserService")
}

func TestUpdateUserController_NotFound(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := mock_user_controller.GetTestingGinContext(recorder)

	id := "507f1f77bcf86cd799439011"
	body := `{"name":"Novo Nome","age":31}`

	mockService := new(mock_user_controller.MockUserService)
	params := gin.Params{gin.Param{Key: "userId", Value: id}}
	mock_user_controller.MakeRequest(
		context,
		params,
		nil,
		"PUT",
		io.NopCloser(strings.NewReader(body)))

	errResp := rest_err.NewBadRequestError("user not found")

	mockService.On("UpdateUserService", id, mock.Anything).Return(errResp)

	controller := NewUserController(mockService)
	controller.UpdateUser(context)

	mockService.AssertExpectations(t)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)

}
