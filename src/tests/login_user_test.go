package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	t.Run("validation_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)

		userLoginRequest := request.UserLoginRequest{
			Email:    "test@mail.com",
			Password: "wrongpassword",
		}

		body, _ := json.Marshal(userLoginRequest)
		mock_user_controller.MakeRequest(
			ctx,
			gin.Params{},
			url.Values{},
			"POST",
			io.NopCloser(bytes.NewBufferString(string(body))))

		UserController.LoginUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})
	t.Run("user_not_found", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)

		userLoginRequest := request.UserLoginRequest{
			Email:    "test@mail.com",
			Password: "Te@st13@",
		}

		body, _ := json.Marshal(userLoginRequest)
		mock_user_controller.MakeRequest(
			ctx,
			gin.Params{},
			url.Values{},
			"POST",
			io.NopCloser(bytes.NewBufferString(string(body))))

		UserController.LoginUser(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	})

}
