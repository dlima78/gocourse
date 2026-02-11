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

var (
	UserDomainInterface model.UserDomainInterface
)

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequest true "User information for registration"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createUser [post]
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("About to start user application",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info,", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserService(domain)
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

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))

}
