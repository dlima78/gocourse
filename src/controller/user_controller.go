package controller

import (
	"github.com/dlima78/gocourse/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserController(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindUserByID(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
