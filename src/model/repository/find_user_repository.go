package repository

import (
	"context"
	"fmt"
	"os"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"github.com/dlima78/gocourse/src/model/repository/entity"
	"github.com/dlima78/gocourse/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository execute successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"),
		zap.String("searchingForId", id))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		errorMessage := fmt.Sprintf("Invalid user ID format: %s", id)
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewBadRequestError(errorMessage)
	}

	logger.Info("ObjectID parsed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("objectIdHex", objectId.Hex()))

	filter := bson.D{{Key: "_id", Value: objectId}}

	logger.Info("Searching with filter",
		zap.String("journey", "findUserByID"),
		zap.String("filter", filter.String()))

	err = collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByID"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmailAndPassword repository",
		zap.String("journey", "findUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User or password invalid"
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
			return nil, rest_err.NewUnauthorizedError(errorMessage)
		}
		errorMessage := "Error trying to find user by email and password"
		logger.Error(errorMessage, err, zap.String("journey", "	findUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository execute successfully",
		zap.String("journey", "findUserByEmailAndPassword"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
