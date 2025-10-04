package config

import (
	"os"
	"github.com/joho/godotenv"
)

type JWTConfig struct {
	SecretKey string
}

func LoadJWTConfig() (*JWTConfig) {
	godotenv.Load()

	return &JWTConfig{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}