package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"projects_tasks/configs"
)

type DatabasePSQL struct {
	DB *sqlx.DB
}

func NewDatabasePSQL(cfg *configs.Config) (*DatabasePSQL, error) {
	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode)

	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	logrus.Info(url)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DatabasePSQL{
		DB: db,
	}, nil
}
