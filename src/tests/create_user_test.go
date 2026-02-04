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

func TestCreateUser(t *testing.T) {
	t.Run("validation_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "wrongmail.com",
			Password: "TEste@123",
			Name:     "João",
			Age:      45,
		}

		body, _ := json.Marshal(userRequest)

		mock_user_controller.MakeRequest(
			ctx, gin.Params{},
			url.Values{},
			"POST",
			io.NopCloser(bytes.NewBufferString(string(body))))

		UserController.CreateUser(ctx)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
