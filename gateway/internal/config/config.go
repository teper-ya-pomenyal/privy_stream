package config

import (
	"log"
	"os"
	"strings"
)

// desktop client (Tauri) sends requests from its webview origin
var tauriOrigins = []string{"tauri://localhost", "http://tauri.localhost", "https://tauri.localhost"}

type Config struct {
	UserServiceAddress         string
	CatalogServiceAddress      string
	CatalogWriteServiceAddress string
	StreamingServiceAddress    string
	PubKeyAddress              string
	TrackStoragePath           string
	Port                       string
	CORSAllowedOrigins         []string
}

func LoadConfig() *Config {
	// web client origins, comma-separated
	var corsOrigins []string
	for _, o := range strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:8081,http://localhost:1420"), ",") {
		if o = strings.TrimSpace(o); o != "" {
			corsOrigins = append(corsOrigins, o)
		}
	}
	corsOrigins = append(corsOrigins, tauriOrigins...)

	return &Config{
		UserServiceAddress:         getEnv("USER_SERVICE_ADDRESS", "localhost:50051"),
		CatalogServiceAddress:      getEnv("CATALOG_SERVICE_ADDRESS", "localhost:50053"),
		CatalogWriteServiceAddress: getEnv("CATALOG_WRITE_SERVICE_ADDRESS", "localhost:50056"),
		StreamingServiceAddress:    getEnv("STREAMING_SERVICE_ADDRESS", "localhost:50054"),
		PubKeyAddress:              getEnv("PUBLIC_KEY_ADDRESS", "keys/public.pem"),
		TrackStoragePath:           getEnv("TRACK_STORAGE_PATH", "data/tracks"),
		Port:                       getEnv("PORT", "8080"),
		CORSAllowedOrigins:         corsOrigins,
	}
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Printf("environment variable %s is not set, using default", key)
		return def
	}
	return v
}
