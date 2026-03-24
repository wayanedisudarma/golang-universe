package services

import (
	"clean-architecture/internal/entity"
	"clean-architecture/internal/model"
	"clean-architecture/internal/repository"
	"context"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(req model.CreateUserRequest) (*entity.User, error)
	GetUser(id string, context context.Context) (*model.GetUserResponse, error)
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func (userService *userServiceImpl) GetUser(id string, context context.Context) (*model.GetUserResponse, error) {

	user, err := userService.userRepository.GetById(id, context)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return &model.GetUserResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
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
