package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	mock_user_controller "github.com/dlima78/gocourse/src/controller/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestDeleteUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := mock_user_controller.GetTestingGinContext(recorder)
		id := bson.NewObjectID().Hex()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		params := gin.Params{gin.Param{Key: "userId", Value: id}}
		mock_user_controller.MakeRequest(ctx, params, url.Values{}, "DELETE", nil)

		UserController.DeleteUser(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}
