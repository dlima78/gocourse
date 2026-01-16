package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

// MockUserService implementa a interface service.UserDomainService apenas para os testes
// métodos que não são usados nos testes retornam zeros/nil para simplificar

func TestDeleteUserController_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("DELETE", "/users/123", nil)
	c.Request = req
	c.Params = gin.Params{gin.Param{Key: "userId", Value: "123"}}

	mockService := new(mock_user_controller.MockUserService)
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

	mockService := new(mock_user_controller.MockUserService)
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
