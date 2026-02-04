package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/controller"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/controller/model/response"
	"github.com/dlima78/gocourse/src/model/repository"
	"github.com/dlima78/gocourse/src/model/service"
	"github.com/dlima78/gocourse/src/tests/connection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

func TestMain(m *testing.M) {
	err := os.Setenv("MONGODB_USER_DB", "test_user")
	if err != nil {
		return
	}

	closeConnection := func() {}
	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserController(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	os.Exit(m.Run())
}
func TestFindUserByEmail(t *testing.T) {
	t.Run("user_not_found_with_this_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mock_user_controller.GetTestingGinContext(recorder)

		params := gin.Params{gin.Param{Key: "userEmail", Value: "teste@teste.com"}}
		mock_user_controller.MakeRequest(context, params, url.Values{}, "GET", nil)

		UserController.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_this_specified_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := bson.NewObjectID().Hex()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "teste@mail.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		params := gin.Params{gin.Param{Key: "userEmail", Value: "teste@mail.com"}}
		mock_user_controller.MakeRequest(ctx, params, url.Values{}, "GET", nil)

		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

		var response response.UserResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, t.Name(), response.Name)
		assert.Equal(t, "teste@mail.com", response.Email)
	})
}

func TestFindUserByID(t *testing.T) {
	t.Run("user_not_found_with_this_ID", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := mock_user_controller.GetTestingGinContext(recorder)
		id := bson.NewObjectID().Hex()

		params := gin.Params{gin.Param{Key: "userId", Value: id}}
		mock_user_controller.MakeRequest(context, params, url.Values{}, "GET", nil)

		UserController.FindUserByID(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_this_specified_ID", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := bson.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "teste@mail.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		params := gin.Params{gin.Param{Key: "userId", Value: id.Hex()}}
		mock_user_controller.MakeRequest(ctx, params, url.Values{}, "GET", nil)

		UserController.FindUserByID(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

		var response response.UserResponse
		err = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Nil(t, err)
		assert.Equal(t, t.Name(), response.Name)
		assert.Equal(t, "teste@mail.com", response.Email)
	})
}
