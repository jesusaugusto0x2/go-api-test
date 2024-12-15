package service

import (
	"context"

	"example.com/go-api-test/ent"
	"example.com/go-api-test/input"
	"example.com/go-api-test/repository"
)

type UserService interface {
	GetUsers(ctx context.Context) ([]*ent.User, error)
	CreateUser(ctx context.Context, input input.CreateUserInput) (*ent.User, error)
	GetUser(ctx context.Context, id int) (*ent.User, error)
	UpdateUser(ctx context.Context, id int, input input.UpdateUserInput) (*ent.User, error) // Nuevo método

}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUsers(ctx context.Context) ([]*ent.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *userService) CreateUser(ctx context.Context, input input.CreateUserInput) (*ent.User, error) {

	// Check if user with the same email already exists
	existingUser, err := s.repo.GetUserByEmail(ctx, input.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	return s.repo.CreateUser(ctx, input)
}

func (s *userService) GetUser(ctx context.Context, id int) (*ent.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, input input.UpdateUserInput) (*ent.User, error) {
	if input.Email != nil {
		existingUser, err := s.repo.GetUserByEmail(ctx, *input.Email)

		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}

		if existingUser != nil && existingUser.ID != id {
			return nil, ErrEmailAlreadyExists
		}
	}

	return s.repo.UpdateUser(ctx, id, input)
}
