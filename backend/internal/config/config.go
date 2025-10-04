package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port     string
	DatabaseURL string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env não encontrado, usando variáveis do sistema")
	}

	return Config{
		Port:       getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
