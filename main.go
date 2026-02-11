package main

import (
	"context"
	"log"

	_ "github.com/dlima78/gocourse/docs"
	docs "github.com/dlima78/gocourse/docs"
	"github.com/dlima78/gocourse/src/configuration/database/mongodb"
	"github.com/dlima78/gocourse/src/controller"
	"github.com/dlima78/gocourse/src/controller/routes"
	"github.com/dlima78/gocourse/src/model/repository"
	"github.com/dlima78/gocourse/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

// @title Gocourse | HunCoding
// @version 1.0
// @description API for crud operations on users
// @host localhost:8081
// @BasePath /
// @schemes http
// @license MIT
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

	docs.SwaggerInfo.BasePath = "/"

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8081/swagger/doc.json")))
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
