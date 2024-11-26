package routes

import (
	controller "github.com/Deatsilence/lawai-go/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine, controller *controller.UserController) {
	// incomingRoutes.POST("/api/users/verifyemail", controller.VerifyEmail())
	incomingRoutes.POST("/api/users/signup", controller.CreateUserHandler)
	// incomingRoutes.POST("/api/users/login", controller.Login())
	// incomingRoutes.POST("/api/users/logout", controller.Logout())
}
