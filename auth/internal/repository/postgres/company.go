package postgres

import (
	"auth/internal/model"
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
