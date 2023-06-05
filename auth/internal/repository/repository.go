package repository

import (
	"auth/configs"
	"auth/internal/model"
	"auth/internal/repository/postgres"
	"auth/internal/repository/redis"
	"github.com/sirupsen/logrus"
)

type IUserRepository interface {
	Create(user model.UserInput) (int, error)
	GetByEmail(string string) (model.User, error)
	GetUsers(page, offset int) ([]model.User, error)
	GetUser(id int) (model.User, error)
}

type ITokenRepository interface {
	Create(token model.RefreshSession) error
	Get(token string) (model.RefreshSession, error)
}

type ICompanyRepository interface {
	Create(company model.CompanyInput) (int, error)
	GetCompanies(page, offset int) ([]model.Company, error)
	GetCompany(id int) (model.Company, error)
}

type Repository struct {
	User    IUserRepository
	Token   ITokenRepository
	Company ICompanyRepository
}

func New(cfg *configs.Config) (*Repository, error) {
	postgresDB, err := postgres.NewDatabasePSQL(cfg)
	if err != nil {
		return nil, err
	}
	logrus.Info("postgres successfully connected")
	redisDB, err := redis.NewRedis(cfg)
	if err != nil {
		return nil, err
	}
	logrus.Info("redis successfully connected")

	user := postgres.NewUser(postgresDB.DB)
	token := redis.NewToken(redisDB.Redis)
	company := postgres.NewCompany(postgresDB.DB)
	return &Repository{
		User:    user,
		Token:   token,
		Company: company,
	}, nil
}
