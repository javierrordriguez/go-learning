package service

import (
	"context"

	"go-learning/internal/model"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
