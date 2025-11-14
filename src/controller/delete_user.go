package controller

import (
	"net/http"

	"github.com/dlima78/gocourse/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

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
