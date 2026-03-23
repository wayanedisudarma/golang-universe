package services

import (
	"clean-architecture/internal/entity"
	"clean-architecture/internal/model"
	"clean-architecture/internal/repository"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(req model.CreateUserRequest) (*entity.User, error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (userService *userServiceImpl) Create(req model.CreateUserRequest) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := userService.userRepository.Create(user); err != nil {
		return nil, err
	}

	slog.Info("User created", "email", user.Email)
	return user, nil
}
