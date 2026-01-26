package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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
