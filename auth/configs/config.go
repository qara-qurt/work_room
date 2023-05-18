package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	SSLMode  string
	DBName   string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "8080"
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		dbHost = "localhost"
	}
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		dbPort = "5432"
	}
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		dbUser = "postgres"
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		dbPassword = "password"
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		dbName = "postgres"
	}
	SSLMode, ok := os.LookupEnv("SSL_MODE")
	if !ok {
		SSLMode = "disable"
	}

	return &Config{
		Server: Server{
			Port: port,
		},
		Database: Database{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			DBName:   dbName,
			SSLMode:  SSLMode,
		},
	}, nil
}
