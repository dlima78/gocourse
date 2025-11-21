package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/model"
	"github.com/dlima78/gocourse/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"

	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user),
		zap.String("journey", "loginUser"),
	)

	userId := c.Param("userId")

	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByID services",
			err,
			zap.String("journey", "findUserByID"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "FindUserByEmail"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user),
		zap.String("journey", "loginUser"),
	)

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail", err, zap.String("journey", "FindUserByEmail"))
		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email")

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error("Error trying to find user by Email", err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUserByEmail controller executed successfully",
		zap.String("journey", "FindUserByEmail"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
