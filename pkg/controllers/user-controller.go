package controllers

import (
	"net/http"

	"github.com/Deatsilence/lawai-go/helpers/base" // base paketini import edin
	models "github.com/Deatsilence/lawai-go/pkg/models/mongo"
	"github.com/Deatsilence/lawai-go/pkg/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		base.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	_, statusCode, err := uc.userService.CreateUser(c.Request.Context(), user)
	if err != nil {
		base.ErrorResponse(c, statusCode, "Failed to create user", err.Error())
		return
	}

	base.SuccessResponse(c, statusCode, "User created successfully", gin.H{})
}
