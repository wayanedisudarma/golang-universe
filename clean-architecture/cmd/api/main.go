package main

import (
	handlers "clean-architecture/internal/api/handlers"
	routes "clean-architecture/internal/api/routes"
	"clean-architecture/internal/config"
	"clean-architecture/internal/repository"
	"clean-architecture/internal/services"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(config.JSONLogger())
	config.InitLogger()
	appConfig, _ := config.NewConfig()
	db := config.NewDatabase(appConfig)
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	routes.NewRouter(routes.RouterConfig{
		UserHandler: userHandler,
	}, router)

	if err := router.Run(); err != nil {
		slog.Error("Failed to start server", err)
		os.Exit(1)
	}
}
