package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userID, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userID),
		zap.String("journey", "updateUser"),
	)

	return nil
}
