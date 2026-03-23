package api

import (
	"clean-architecture/internal/model"
	"clean-architecture/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) GetUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"user": "John Doe"})
}

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	var req model.CreateUserRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error()))
		return
	}

	user, err := userHandler.userService.Create(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest("Could not create user"))
		return
	}

	context.JSON(http.StatusCreated, model.ResponseOkWithData(user))
}
