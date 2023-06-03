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

type ICompany interface {
	Create(input model.CompanyInput) (int, error)
}

type Service struct {
	User    IUser
	Company ICompany
}

func New(repo *repository.Repository, cfg *configs.Config) *Service {
	user := NewUser(repo.User, repo.Token, cfg.Server.HMACSecret)
	company := NewCompany(repo.Company)
	return &Service{
		User:    user,
		Company: company,
	}
}
