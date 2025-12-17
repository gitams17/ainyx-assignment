package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err!= nil {
		log.Println("No.env file found, using default env variables")
	}

	return Config{
		DBUrl: getEnv("DATABASE_URL", "postgresql://user:password@localhost:5432/fiber_db?sslmode=disable"),
		Port:  getEnv("PORT", ":3000"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}