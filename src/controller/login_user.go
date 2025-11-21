package controller

import (
	"net/http"

	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/configuration/validation"
	"github.com/dlima78/gocourse/src/controller/model/request"
	"github.com/dlima78/gocourse/src/model"
	"github.com/dlima78/gocourse/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init login user application",
		zap.String("journey", "loginUser"),
	)
	var userLoginRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		logger.Error("Error trying to validate user info,", err,
			zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)

	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
