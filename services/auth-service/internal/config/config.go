package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
	Port      string
}

func LoadEnv() *Config {
	_ = godotenv.Load()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8087"
	}

	return &Config{
		JWTSecret: secret,
		Port:      port,
	}
}
