package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Port        string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	cfg := Config{}

	cfg.Environment = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	return cfg
}
