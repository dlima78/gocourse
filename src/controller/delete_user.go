package controller

import (
	"net/http"

	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteUser/{userId} [delete]
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)
	userId := c.Param("userId")

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error(
			"Error trying to call DeleteUser service",
			err,
			zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"DeleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}
