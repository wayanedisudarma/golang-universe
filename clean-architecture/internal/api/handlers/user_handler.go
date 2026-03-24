package api

import (
	"clean-architecture/internal/model"
	"clean-architecture/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (userHandler *UserHandler) GetUser(context *gin.Context) {
	userId := context.Param("userId")

	if _, err := uuid.Parse(userId); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest("Invalid user id"))
		return
	}

	getUserResponse, err := userHandler.userService.GetUser(userId, context.Request.Context())
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error()))
		return
	}
	context.JSON(http.StatusOK, model.ResponseOkWithData(getUserResponse))
}

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	var req model.CreateUserRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error()))
		return
	}

	user, err := userHandler.userService.Create(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error()))
		return
	}

	context.JSON(http.StatusCreated, model.ResponseOkWithData(user))
}
