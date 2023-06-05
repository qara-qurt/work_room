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
	GetUsers(page, offset int) ([]model.User, error)
	GetUser(id int) (model.User, error)
}

type ICompany interface {
	Create(input model.CompanyInput) (int, error)
	GetCompanies(page, offset int) ([]model.Company, error)
	GetCompany(id int) (model.Company, error)
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
