package service

import (
	"auth/internal/model"
	"auth/internal/repository"
)

type IUser interface {
	Create(user model.UserInput) (int, error)
}

type Service struct {
	User IUser
}

func New(repo *repository.Repository) *Service {
	user := NewUser(repo.User)
	return &Service{
		User: user,
	}
}
