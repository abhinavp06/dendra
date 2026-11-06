package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Environment string
	Port        string
}

func loadConfig() config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	cfg := config{}

	cfg.Environment = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	return cfg
}
