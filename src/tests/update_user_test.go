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

func TestUpdateUser(t *testing.T) {
	t.Run("validation_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := "wrong-id"

		userRequest := request.UserRequest{
			Email:    "test@mail.com",
			Password: "TEste@123",
			Name:     "João",
			Age:      45,
		}

		body, _ := json.Marshal(userRequest)

		params := gin.Params{gin.Param{Key: "userId", Value: id}}
		mock_user_controller.MakeRequest(
			ctx,
			params,
			url.Values{},
			"PUT",
			io.NopCloser(bytes.NewBufferString(string(body))))

		UserController.UpdateUser(ctx)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
