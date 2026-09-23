package config

import (
	"log"
	"os"
)

type Config struct {
	TrackStoragePath           string
	CatalogServiceAddress      string
	CatalogWriteServiceAddress string
	Port                       string
}

func NewConfig() *Config {
	tsp := os.Getenv("TRACK_STORAGE_PATH")
	if tsp == "" {
		log.Fatal("environment variable missing: TRACK_STORAGE_PATH")
	}
	csa := os.Getenv("CATALOG_SERVICE_ADDRESS")

	if tsp == "" {
		log.Fatal("environment variable missing: CATALOG_SERVICE_ADDRESS")
	}
	cwsa := os.Getenv("CATALOG_WRITE_SERVICE_ADDRESS")
	if cwsa == "" {
		log.Fatal("environment variable missing: CATALOG_WRITE_SERVICE_ADDRESS")
	}
	port := os.Getenv("STREAMING_SERVICE_PORT")
	if tsp == "" {
		log.Fatal("environment variable missing: STREAMING_SERVICE_PORT")
	}
	return &Config{
		TrackStoragePath:           tsp,
		CatalogServiceAddress:      csa,
		CatalogWriteServiceAddress: cwsa,
		Port:                       port,
	}
}
