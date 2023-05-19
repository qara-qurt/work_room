package service

import (
	"auth/configs"
	"auth/internal/model"
	"auth/internal/repository"
)

type IUser interface {
	Create(user model.UserInput) (int, error)
	SignIn(user model.UserAuthInput) (string, string, error)
	RefreshTokens(refreshToken string) (string, string, error)
}

type Service struct {
	User IUser
}

func New(repo *repository.Repository, cfg *configs.Config) *Service {
	user := NewUser(repo.User, repo.Token, cfg.Server.HMACSecret)
	return &Service{
		User: user,
	}
}
