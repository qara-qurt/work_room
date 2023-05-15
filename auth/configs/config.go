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

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	dbHost, ok := os.LookupEnv("DBHOST")
	if !ok {
		dbHost = "localhost"
	}
	dbPort, ok := os.LookupEnv("DBPORT")
	if !ok {
		dbPort = "5432"
	}
	dbUser, ok := os.LookupEnv("DBUSER")
	if !ok {
		dbUser = "postgres"
	}
	dbPassword, ok := os.LookupEnv("DBPASSWORD")
	if !ok {
		dbPassword = "password"
	}
	dbName, ok := os.LookupEnv("DBNAME")
	if !ok {
		dbName = "postgres"
	}
	dbSSLMode, ok := os.LookupEnv("DBSSLMODE")
	if !ok {
		dbSSLMode = "disable"
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
			SSLMode:  dbSSLMode,
		},
	}, nil
}
