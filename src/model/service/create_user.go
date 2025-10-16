package service

import (
	"fmt"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(ud)
	return nil
}
