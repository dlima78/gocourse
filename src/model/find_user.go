package model

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser(userID string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init find user model", zap.String("journey", "findUser"))

	return nil, nil
}
