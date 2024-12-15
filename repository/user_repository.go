package repository

import (
	"context"

	"example.com/go-api-test/ent"
	"example.com/go-api-test/ent/user"
	"example.com/go-api-test/input"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	CreateUser(ctx context.Context, input input.CreateUserInput) (*ent.User, error)
	GetUserByID(ctx context.Context, id int) (*ent.User, error)
	UpdateUser(ctx context.Context, id int, input input.UpdateUserInput) (*ent.User, error)
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

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	user, err := r.client.User.Query().Where(user.EmailEQ(email)).Only(ctx)

	// Prevent error throwing when user is not found, just return nil
	if ent.IsNotFound(err) {
		return nil, nil
	}

	return user, err
}

func (r *userRepository) CreateUser(ctx context.Context, input input.CreateUserInput) (*ent.User, error) {
	return r.client.User.Create().
		SetName(input.Name).
		SetEmail(input.Email).
		Save(ctx)
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.client.User.Query().Where(user.IDEQ(id)).Only(ctx)
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, input input.UpdateUserInput) (*ent.User, error) {
	update := r.client.User.UpdateOneID(id)

	if input.Name != nil {
		update.SetName(*input.Name)
	}
	if input.Email != nil {
		update.SetEmail(*input.Email)
	}

	return update.Save(ctx)
}
