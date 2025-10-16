package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userID string) *rest_err.RestErr {
	logger.Info("Init delete user model", zap.String("journey", "deleteUser"))

	return nil
}
