package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	godotenv.Load()
}

func AppServer() *App {
	return &App{
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),
	}
}

func GetPostgresConfig() *PostgresDB {
	return &PostgresDB{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USERNAME"),
		DBName:   os.Getenv("POSTGRES_DATABASE"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func NewLogger() *zap.Logger {
	var (
		logger *zap.Logger
		err    error
		env    string
	)

	env = os.Getenv("APP_ENV")

	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic(err)
	}

	return logger
}

func GetPostgresSync() *PostgresSync {

	return &PostgresSync{}
}
