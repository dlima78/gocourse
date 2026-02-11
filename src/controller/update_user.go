package controller

import (
	"net/http"
	"strings"

	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/dlima78/gocourse/src/configuration/validation"
	"github.com/dlima78/gocourse/src/controller/model/request"
	"github.com/dlima78/gocourse/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateUser updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateUser/{userId} [put]
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("About to start user application",
		zap.String("journey", "updateUser"),
	)
	var userUpdateRequest request.UserUpdateRequest

	userId := c.Param("userId")

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil ||
		strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user info,", err,
			zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userUpdateRequest.Name,
		userUpdateRequest.Age,
	)

	err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to call UpdateUser service",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
