package repository

import (
	"auth/configs"
	"auth/internal/model"
	"auth/internal/repository/postgres"
	"auth/internal/repository/redis"
)

type IUserRepository interface {
	Create(user model.UserInput) (int, error)
}

type Repository struct {
	User IUserRepository
}

func New(cfg *configs.Config) (*Repository, error) {
	postgresDB, err := postgres.NewDatabasePSQL(cfg)
	if err != nil {
		return nil, err
	}
	redisDB, err := redis.NewRedis(cfg)
	if err != nil {
		return nil, err
	}

	user := postgres.NewUser(postgresDB.DB, redisDB.Redis)
	return &Repository{
		User: user,
	}, nil
}
