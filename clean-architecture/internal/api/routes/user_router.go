package api

import (
	handlers "clean-architecture/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func MapUserRoutes(group *gin.RouterGroup, userHandler *handlers.UserHandler) {
	group.GET("/:userId", userHandler.GetUser)
	group.POST("/", userHandler.CreateUser)
}
