package config

import (
	"log"
	"os"
)

type Config struct {
	UserServiceAddress    string
	CatalogServiceAddress string
	PubKeyAddress         string
	Port                  string
}

func LoadConfig() *Config {
	usAddress := os.Getenv("USER_SERVICE_ADDRESS")
	if usAddress == "" {
		log.Fatal("environment variable missing: USER_SERVICE_ADDRESS")
	}

	catalogAddress := os.Getenv("CATALOG_SERVICE_ADDRESS")
	if catalogAddress == "" {
		log.Fatal("environment variable missing: CATALOG_SERVICE_ADDRESS")
	}

	pka := os.Getenv("PUBLIC_KEY_ADDRESS")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		UserServiceAddress:    usAddress,
		CatalogServiceAddress: catalogAddress,
		PubKeyAddress:         pka,
		Port:                  port,
	}
}
