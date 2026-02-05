package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/controller/model/request"
	"github.com/dlima78/gocourse/src/model/repository/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUpdateUser(t *testing.T) {
	t.Run("validation_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := "wrong-id"

		userRequest := request.UserRequest{
			Email:    "wrongmail.com",
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

	t.Run("success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := bson.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{
				"_id":   id,
				"email": "test@mail.com",
				"name":  "Eduardo",
				"age":   47,
			})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{{Key: "userId", Value: id.Hex()}}

		userUpdateRequest := request.UserUpdateRequest{
			Name: "João",
			Age:  25,
		}

		b, _ := json.Marshal(userUpdateRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		mock_user_controller.MakeRequest(
			ctx,
			param,
			url.Values{},
			"PUT",
			stringReader,
		)

		UserController.UpdateUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

		userEntity := entity.UserEntity{}
		filter := bson.D{{Key: "_id", Value: id}}
		_ = Database.
			Collection("test_user").
			FindOne(context.Background(), filter).Decode(&userEntity)

		assert.EqualValues(t, userUpdateRequest.Name, userEntity.Name)
		assert.EqualValues(t, userUpdateRequest.Age, userEntity.Age)
	})
}
