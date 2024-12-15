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
	return s.repo.CreateUser(ctx, input)
}
