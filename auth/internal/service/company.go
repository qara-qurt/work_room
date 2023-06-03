package service

import (
	"auth/internal/model"
	"auth/internal/repository"
)

type Company struct {
	repo repository.ICompanyRepository
}

func NewCompany(repo repository.ICompanyRepository) *Company {
	return &Company{
		repo: repo,
	}
}

func (c *Company) Create(company model.CompanyInput) (int, error) {
	return c.repo.Create(company)
}
