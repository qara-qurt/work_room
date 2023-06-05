package model

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
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
	BirthDate string `json:"birth_date" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	Email     string `json:"email" validate:"email,gte=4"`
	Phone     string `json:"phone" validate:"e164"`
	Password  string `json:"password" validate:"required,gte=6"`
	Role      string `json:"role" validate:"required,gte=4"`
}

func (u *UserInput) Validate() error {
	return validate.Struct(u)
}

type UserAuthInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

func (u *UserAuthInput) Validate() error {
	return validate.Struct(u)
}

type User struct {
	ID        uint       `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Surname   string     `json:"surname" db:"surname"`
	Img       *string    `json:"img" db:"img"`               //can be null
	Position  *string    `json:"position" db:"position"`     //can be null
	Location  *string    `json:"location" db:"location"`     //can be null
	BirthDate *time.Time `json:"birth_date" db:"birth_date"` //can be null
	Gender    string     `json:"gender" db:"gender"`
	Role      string     `json:"role" db:"role"`
	Email     string     `json:"email" db:"email"`
	Phone     *string    `json:"phone" db:"phone"` //can be null
	Password  string     `json:"password" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

var ErrorUserAlreadyExist = errors.New("user with this email already exist")
var ErrorUserPhoneAlreadyLinked = errors.New("this phone number is already linked to another user")
var ErrorUserNotFound = errors.New("user with this email not found")
var ErrorUserPassword = errors.New("user password is not correct")
var ErrorUserIdNotFound = errors.New("user with this id doesnt exist")
