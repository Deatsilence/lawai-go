package main

import (
	"log"
	"os"

	"github.com/Deatsilence/lawai-go/database"
	businesslogic "github.com/Deatsilence/lawai-go/pkg/business-logic"
	"github.com/Deatsilence/lawai-go/pkg/controllers"
	"github.com/Deatsilence/lawai-go/pkg/repositories"
	"github.com/Deatsilence/lawai-go/pkg/services"
	"github.com/Deatsilence/lawai-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// github.com/Deatsilence/lawai-go.git

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file `%v`", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// MongoDB koleksiyonunu oluştur
	userCollection := database.OpenCollection(database.Client, "users")
	userRepo := repositories.NewUserRepository(userCollection)

	// Service ve Controller'ı başlat
	userBL := businesslogic.NewUserBL()
	commonBL := businesslogic.NewCommonBL()
	userService := services.NewUserService(userRepo, userBL, commonBL)
	userController := controllers.NewUserController(userService)

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router, userController)

	router.Run(":" + port)
}
