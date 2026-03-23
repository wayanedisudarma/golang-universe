package main

import (
	handlers "clean-architecture/internal/api/handlers"
	routes "clean-architecture/internal/api/routes"
	"clean-architecture/internal/config"
	"clean-architecture/internal/repository"
	"clean-architecture/internal/services"
	"log"
)

func main() {

	db := config.NewDatabase()
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	router := routes.NewRouter(routes.RouterConfig{
		UserHandler: userHandler,
	})

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
