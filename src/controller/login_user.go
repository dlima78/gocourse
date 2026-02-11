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

// LoginUser allows a user to log in and obtain an authentication token.
// @Summary User Login
// @Description Allows a user to log in and receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userLogin body request.UserLoginRequest true "User login credentials"
// @Success 200 {object} response.UserResponse "Login successful, authentication token provided"
// @Header 200 {string} Authorization "Authentication token"
// @Failure 403 {object} rest_err.RestErr "Error: Invalid login credentials"
// @Router /login [post]
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
			"Error trying to call Login service",
			err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"LoginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
