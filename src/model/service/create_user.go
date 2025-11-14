package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	user, _ := ud.userRepository.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		logger.Error("Email already exists", nil, zap.String("journey", "createUser"))
		return nil, rest_err.NewBadRequestError("Email already exists")
	}

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"createUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)

	return userDomainRepository, nil
}
