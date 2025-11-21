package service

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init login user model", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordService(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)

	return user, token, nil
}
