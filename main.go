// ...existing code...
package main

import (
	"context"
	"log"

	_ "github.com/dlima78/gocourse/docs"
	"github.com/dlima78/gocourse/src/configuration/database/mongodb"
	"github.com/dlima78/gocourse/src/controller"
	"github.com/dlima78/gocourse/src/controller/routes"
	"github.com/dlima78/gocourse/src/model/repository"
	"github.com/dlima78/gocourse/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect database, error=%s\n", err)
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserController(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
