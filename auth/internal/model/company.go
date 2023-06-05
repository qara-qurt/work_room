package model

import "errors"

type CompanyInput struct {
	Name        string `json:"name" validate:"required,gte=3"`
	OwnerID     int    `json:"owner_id" validate:"required"`
	Description string `json:"description"`
}

func (c *CompanyInput) Validate() error {
	return validate.Struct(c)
}

type Company struct {
	ID          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	OwnerID     int    `json:"owner_id" db:"owner_id"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

var ErrorCompanyAlreadyExist = errors.New("company with this name already exist")
var ErrorCompanyIdNotFound = errors.New("company with this id doesnt exist")
