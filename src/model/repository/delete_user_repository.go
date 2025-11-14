package repository

import (
	"context"
	"fmt"
	"os"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUser"),
		zap.String("searchingForId", userId))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	objectId, err := bson.ObjectIDFromHex(userId)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid user ID format: %s", userId)
		logger.Error(errorMessage, err, zap.String("journey", "deleteUser"))
		return rest_err.NewBadRequestError(errorMessage)
	}

	filter := bson.D{{Key: "_id", Value: objectId}}

	logger.Info("Searching with filter",
		zap.String("journey", "deleteUser"),
		zap.String("filter", filter.String()))

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", userId)
			logger.Error(errorMessage, err, zap.String("journey", "deleteUser"))
			return rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to delete user from database"
		logger.Error(errorMessage, err, zap.String("journey", "deleteUser"))

		return rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("DeleteUser repository executed successfully", zap.String("journey", "deleteUser"))

	return nil
}
