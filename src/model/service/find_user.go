package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailService(email string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init find findUserByEmailService service", zap.String("journey", "findUserByEmailService"))

	return ud.userRepository.FindUserByEmail(email)
}
func (ud *userDomainService) FindUserByIDService(id string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init find findUserByIDService service", zap.String("journey", "findUserByIDService"))

	return ud.userRepository.FindUserByID(id)
}
