package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("journey", "updateUser"))

	return nil
}
