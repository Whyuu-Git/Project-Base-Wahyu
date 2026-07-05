package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}


func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		AppPort: getEnv("APP_PORT", "14044"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "1234"),
		DBName:     getEnv("DB_NAME", "project-base-wahyu"),
	}
}


func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
