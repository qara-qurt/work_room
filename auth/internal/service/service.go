package service

import "auth/internal/repository"

type IUser interface{}

type Service struct {
	User IUser
}

func New(repo *repository.Repository) *Service {
	user := NewUser(repo.User)
	return &Service{
		User: user,
	}
}
