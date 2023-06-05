package postgres

import (
	"auth/internal/model"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Company struct {
	db *sqlx.DB
}

func NewCompany(db *sqlx.DB) *Company {
	return &Company{
		db: db,
	}
}

func (c *Company) Create(company model.CompanyInput) (int, error) {
	var companyId int
	query := `INSERT INTO company (name,owner_id,description) VALUES ($1,$2,$3) RETURNING id`

	err := c.db.QueryRowx(query, company.Name, company.OwnerID, company.Description).Scan(&companyId)
	if err != nil {
		if err.(*pq.Error).Constraint == "company_name_key" {
			return 0, model.ErrorCompanyAlreadyExist
		}
		return 0, err
	}

	return companyId, nil
}

func (u *Company) GetCompanies(page, offset int) ([]model.Company, error) {
	query := fmt.Sprintf("SELECT * FROM company LIMIT %d OFFSET %d", offset, (page-1)*offset)
	var companies []model.Company

	if err := u.db.Select(&companies, query); err != nil {
		return companies, err
	}

	return companies, nil
}

func (u *Company) GetCompany(id int) (model.Company, error) {
	query := "SELECT * FROM company WHERE id = $1"
	var company model.Company

	err := u.db.Get(&company, query, id)
	if err == sql.ErrNoRows {
		return company, model.ErrorCompanyIdNotFound
	} else if err != nil {
		return company, err
	}

	return company, nil
}
