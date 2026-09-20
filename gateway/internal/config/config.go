package config

import (
	"log"
	"os"
)

type Config struct {
	UserServiceAddress         string
	CatalogServiceAddress      string
	CatalogWriteServiceAddress string
	StreamingServiceAddress    string
	PubKeyAddress              string
	TrackStoragePath           string
	Port                       string
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
	catalogWriteAddress := os.Getenv("CATALOG_WRITE_SERVICE_ADDRESS")
	if catalogWriteAddress == "" {
		log.Fatal("environment variable missing: CATALOG_WRITE_SERVICE_ADDRESS")
	}
	streamingAddress := os.Getenv("STREAMING_SERVICE_ADDRESS")
	if streamingAddress == "" {
		log.Fatal("environment variable missing: CATALOG_SERVICE_ADDRESS")
	}

	pka := os.Getenv("PUBLIC_KEY_ADDRESS")

	trackStoragePath := os.Getenv("TRACK_STORAGE_PATH")
	if trackStoragePath == "" {
		log.Fatal("environment variable missing: TRACK_STORAGE_PATH")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		UserServiceAddress:         usAddress,
		CatalogServiceAddress:      catalogAddress,
		CatalogWriteServiceAddress: catalogWriteAddress,
		StreamingServiceAddress:    streamingAddress,
		PubKeyAddress:              pka,
		TrackStoragePath:           trackStoragePath,
		Port:                       port,
	}
}
