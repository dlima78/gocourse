package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/dlima78/gocourse/src/controller"
	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/dlima78/gocourse/src/model/repository"
	"github.com/dlima78/gocourse/src/model/service"
	"github.com/dlima78/gocourse/src/tests/connection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
}
