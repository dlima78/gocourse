package repository

import (
	"context"
	"fmt"
	"os"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"github.com/dlima78/gocourse/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser repository",
		zap.String("journey", "updateUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	fmt.Println("VALUE: ", value)

	objectId, err := rest_err.NewObjectIDFromHex(userId)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid user ID format: %s", userId)
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return rest_err.NewBadRequestError(errorMessage)
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: value}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user",
			err,
			zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User updated successfully",
		zap.String("journey", "updateUser"),
		zap.String("matchedCount", fmt.Sprintf("%v", result.MatchedCount)),
		zap.String("modifiedCount", fmt.Sprintf("%v", result.ModifiedCount)),
	)

	return nil
}
