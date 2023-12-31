package config

import (
	"errors"

	"github.com/joho/godotenv"
)

func LoadEnvVars(env string) error {
	var err error

	switch env {

	case "dev", "development":
		err = godotenv.Load("./config/.env.dev")

	case "prod", "production":
		err = godotenv.Load("./config/env.prod")

	default:
		err = errors.New("no env vars found")

	}

	return err
}
