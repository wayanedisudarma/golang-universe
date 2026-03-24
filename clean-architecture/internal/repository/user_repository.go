package repository

import (
	"clean-architecture/internal/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetById(id string, context context.Context) (*entity.User, error)
}

type pgUserRepository struct {
	db *gorm.DB
}

func (r *pgUserRepository) GetById(id string, context context.Context) (*entity.User, error) {
	var user entity.User

	result := r.db.WithContext(context).First(&user, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
