package config

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host string
	Port int
	User string
	Password string
	DBName string
	SSLMode string
}

func LoadDBConfig() *DBConfig {
    godotenv.Load()

    port, err := strconv.Atoi(os.Getenv("DB_PORT"))
    if err != nil || port == 0 {
        port = 5432
    }

    return &DBConfig{
        Host:     getEnv("DB_HOST", "localhost"),
        Port:     port,
        User:     getEnv("DB_USER", "laundry_user"),
        Password: getEnv("DB_PASSWORD", "mypassword"),
        DBName:   getEnv("DB_NAME", "laundry_db"),
        SSLMode:  getEnv("DB_SSLMODE", "disable"),
    }
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}