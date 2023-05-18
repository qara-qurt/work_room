package configs

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Server   Server
	Database Database
	Redis    Redis
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

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: Server{
			Port: GetEnv("SERVER_PORT", "8080"),
		},
		Database: Database{
			Host:     GetEnv("DB_HOST", "localhost"),
			Port:     GetEnv("DB_PORT", "5432"),
			User:     GetEnv("DB_USER", "postgres"),
			Password: GetEnv("DB_PASSWORD", "password"),
			DBName:   GetEnv("DB_NAME", "postgres"),
			SSLMode:  GetEnv("SSL_MODE", "disable"),
		},
		Redis: Redis{
			Host:     GetEnv("REDIS_HOST", "localhost"),
			Port:     GetEnv("REDIS_PORT", "6379"),
			Password: GetEnv("REDIS_PASSWORD", ""),
			DB: func() int {
				dbStr := GetEnv("REDIS_DB", "0")
				db, _ := strconv.Atoi(dbStr)
				return db
			}(),
		},
	}, nil
}
