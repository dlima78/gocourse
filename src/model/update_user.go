package model

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser(userID string) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("journey", "updateUser"))

	return nil
}
