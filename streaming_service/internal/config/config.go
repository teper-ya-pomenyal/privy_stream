package config

import "os"

type Config struct {
	TrackStoragePath      string
	CatalogServiceAddress string
	Port                  string
}

func NewConfig() *Config {
	tsp := os.Getenv("TRACK_STORAGE_PATH")
	csa := os.Getenv("CATALOG_SERVICE_ADDRESS")
	port := os.Getenv("STREAMING_SERVICE_PORT")
	return &Config{
		TrackStoragePath:      tsp,
		CatalogServiceAddress: csa,
		Port:                  port,
	}
}
