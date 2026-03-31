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
	lang := context.GetString("lang")
	traceId := context.GetString("traceId")
	userId := context.Param("userId")

	if _, err := uuid.Parse(userId); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest("INVALID_USER_ID", lang, traceId))
		return
	}

	getUserResponse, err := userHandler.userService.GetUser(userId, context.Request.Context())
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error(), lang, traceId))
		return
	}

	context.JSON(http.StatusOK, model.ResponseOkWithData(getUserResponse))
}

func (userHandler *UserHandler) CreateUser(context *gin.Context) {
	lang := context.GetString("lang")
	traceId := context.GetString("traceId")
	var req model.CreateUserRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error(), lang, traceId))
		return
	}

	user, err := userHandler.userService.Create(req)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.ResponseBadRequest(err.Error(), lang, traceId))
		return
	}

	context.JSON(http.StatusCreated, model.ResponseOkWithData(user))
}
