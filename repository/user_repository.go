package repository

import (
	"context"

	"example.com/go-api-test/ent"
	"example.com/go-api-test/ent/user"
	"example.com/go-api-test/input"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*ent.User, error)
	CreateUser(ctx context.Context, input input.CreateUserInput) (*ent.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().Order(ent.Asc(user.FieldID)).All(ctx)
}

func (r *userRepository) CreateUser(ctx context.Context, input input.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().
		SetName(input.Name).
		SetEmail(input.Email).
		Save(ctx)
}
