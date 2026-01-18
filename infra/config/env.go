package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseConfig *DatabaseConfig
}

func LoadEnv() *Env {
	err := godotenv.Load()

	getEnv := func(key string) string {
		return os.Getenv(key)
	}

	if err != nil {
		panic(err)
	}

	return &Env{
		DatabaseConfig: &DatabaseConfig{
			Host:     getEnv("DATABASE_HOST"),
			Port:     getEnv("DATABASE_PORT"),
			User:     getEnv("DATABASE_USER"),
			Password: getEnv("DATABASE_PASSWORD"),
			Name:     getEnv("DATABASE_NAME"),
		},
	}
}
