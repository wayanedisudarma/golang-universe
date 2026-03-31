package main

import (
	handlers "clean-architecture/internal/api/handlers"
	routes "clean-architecture/internal/api/routes"
	"clean-architecture/internal/config"
	"clean-architecture/internal/i18n"
	"clean-architecture/internal/middleware"
	"clean-architecture/internal/repository"
	"clean-architecture/internal/services"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	i18n.Init()
	router := gin.New()
	router.Use(middleware.Middleware())
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
		slog.Error("Failed to start server", "details", err.Error())
		os.Exit(1)
	}
}
