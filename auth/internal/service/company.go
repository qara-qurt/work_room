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

func (c *Company) GetCompanies(page, offset int) ([]model.Company, error) {
	return c.repo.GetCompanies(page, offset)
}

func (c *Company) GetCompany(id int) (model.Company, error) {
	return c.repo.GetCompany(id)
}
