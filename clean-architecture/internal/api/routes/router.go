package api

import (
	handlers "clean-architecture/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserHandler *handlers.UserHandler
}

func NewRouter(config RouterConfig) *gin.Engine {
	router := gin.Default()

	MapUserRoutes(router.Group("/api/v1/user"), config.UserHandler)

	return router
}
