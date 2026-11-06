package config

import (
	"os"
	"strconv"
)

type Config struct {
	Environment string
	Port        int
}

func Load() (Config, error) {
	cfg := Config{}

	cfg.Environment = os.Getenv("ENV", )
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return Config{}, err
	}

	cfg.Port = port

	return cfg, nil
}
