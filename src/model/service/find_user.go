package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(userID string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init find user model", zap.String("journey", "findUser"))

	return nil, nil
}
