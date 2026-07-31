package user

import (
	"context"
	"errors"
)

type MemUserRepository struct {
	users map[int64]*User
}

func NewUserRepository() *MemUserRepository {
	return &MemUserRepository{
		users: map[int64]*User{
			1: {
				ID:    1,
				Name:  "Javier",
				Email: "javier@example.com",
			},
		},
	}
}

func (r *MemUserRepository) GetByID(ctx context.Context, id int64) (*User, error) {
	user, ok := r.users[id]

	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *MemUserRepository) Create(ctx context.Context, user *User) error {
	r.users[user.ID] = user
	return nil
}
