package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userID string) *rest_err.RestErr {
	logger.Info("Init delete user model", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userID)
	if err != nil {
		logger.Error(
			"Error trying to delete user from database",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("Delete user model executed successfully", zap.String("journey", "deleteUser"))

	return nil
}
