package controller

import (
	rest_err "github.com/dlima78/gocourse/src/configuration"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamou a rota errada")
	c.JSON(err.Code, err)
}
