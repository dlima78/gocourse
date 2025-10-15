package model

import (
	"fmt"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
