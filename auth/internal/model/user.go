package model

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type UserInput struct {
	Name      string `json:"name" validate:"required,gte=2"`
	Surname   string `json:"surname" validate:"required,gte=2"`
	Position  string `json:"position" validate:"required"`
	Location  string `json:"location" validate:"required,gte=3"`
	BirthDate string `json:"birth_date" validate:"required" validate:"datetime"`
	Gender    string `json:"gender" validate:"required"`
	Email     string `json:"email" validate:"email,gte=4"`
	Phone     string `json:"phone" validate:"e164"`
	Password  string `json:"password" validate:"required,gte=6"`
	Role      string `json:"role" validate:"required,gte=4"`
}

func (u *UserInput) Validate() error {
	return validate.Struct(u)
}

var ErrorAlreadyExist = errors.New("user with this email already exist")
var ErrorPhoneAlreadyLinked = errors.New("this phone number is already linked to another user")
