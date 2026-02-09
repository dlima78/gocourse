package test

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// encryptPassword encripta a senha usando MD5 (mesma lógica do modelo)
func encryptPassword(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

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

	t.Run("success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := bson.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(),
				bson.M{
					"_id":      id,
					"email":    "test@mail.com",
					"password": encryptPassword("Te@st13@"),
					"name":     "Eduardo",
					"age":      47,
				})

		if err != nil {
			t.Fatal(err)
			return
		}

		userLoginRequest := request.UserLoginRequest{
			Email:    "test@mail.com",
			Password: "Te@st13@",
		}

		body, _ := json.Marshal(userLoginRequest)

		fmt.Println("BODY: ", string(body))
		mock_user_controller.MakeRequest(
			ctx,
			gin.Params{},
			url.Values{},
			"POST",
			io.NopCloser(bytes.NewBufferString(string(body))))

		UserController.LoginUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})

}
