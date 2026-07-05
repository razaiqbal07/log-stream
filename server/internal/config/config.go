package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DATABASE string
	PORT     string
}

func ResolveConfig() *AppConfig {
	godotenv.Load()
	db := os.Getenv("DATABASE")
	return &AppConfig{
		DATABASE: db,
		PORT:     os.Getenv("PORT"),
	}
}
