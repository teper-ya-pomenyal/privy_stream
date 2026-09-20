package config

import (
	"log"
	"os"
)

type Config struct {
	Port      string
	WritePort string
	DSN       string
}

func LoadConfig() *Config {
	return &Config{
		Port:      mustGetEnv("CATALOG_SERVICE_PORT"),
		WritePort: mustGetEnv("CATALOG_WRITE_SERVICE_PORT"),
		DSN:       mustGetEnv("CATALOG_DB_DSN"),
	}
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return v
}
