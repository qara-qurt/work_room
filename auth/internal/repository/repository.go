package repository

import (
	"auth/configs"
)

type IUserRepository interface {
}

type Repository struct {
	User IUserRepository
}

func New(cfg *configs.Config) (*Repository, error) {
	//db, err := postgres.NewDatabasePSQL(cfg)
	//if err != nil {
	//	return nil, err
	//}
	//
	//user := postgres.NewUser(db.DB)
	//return Repository{
	//	User: user,
	//}, nil
	return &Repository{}, nil
}
