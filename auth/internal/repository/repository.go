package repository

import (
	"auth/configs"
	"auth/internal/model"
	"auth/internal/repository/postgres"
)

type IUserRepository interface {
	Create(user model.UserInput) (int, error)
}

type Repository struct {
	User IUserRepository
}

func New(cfg *configs.Config) (*Repository, error) {
	db, err := postgres.NewDatabasePSQL(cfg)
	if err != nil {
		return nil, err
	}

	user := postgres.NewUser(db.DB)
	return &Repository{
		User: user,
	}, nil
}
