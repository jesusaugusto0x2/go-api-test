package config

import (
	"fmt"
	"os"
)

type Config struct {
	DSN string
}

func NewConfig() (*Config, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	vars := []string{user, pass, host, port, name}

	for _, v := range vars {
		if v == "" {
			return nil, fmt.Errorf("missing required environment variables for database configuration")
		}
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, pass, host, port, name,
	)

	return &Config{DSN: dsn}, nil
}
