package memory

import (
	"context"
	"errors"

	"go-learning/internal/model"
)

type UserRepository struct {
	users map[int64]*model.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[int64]*model.User{
			1: {
				ID:    1,
				Name:  "Javier",
				Email: "javier@example.com",
			},
		},
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user, ok := r.users[id]

	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	r.users[user.ID] = user
	return nil
}