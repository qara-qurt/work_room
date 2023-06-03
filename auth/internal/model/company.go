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

var ErrorCompanyAlreadyExist = errors.New("company with this name already exist")
