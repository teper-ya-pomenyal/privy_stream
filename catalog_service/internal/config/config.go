package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port      string
	WritePort string
	DB        DBConfig
}

type DBConfig struct {
	DSN               string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectTimeout    time.Duration
}

func LoadConfig() *Config {
	return &Config{
		Port:      getEnv("CATALOG_SERVICE_PORT", "50053"),
		WritePort: getEnv("CATALOG_WRITE_SERVICE_PORT", "50056"),
		DB: DBConfig{
			DSN:               getEnv("CATALOG_DB_DSN", "postgres://postgres:postgres@localhost:5432/catalog_db?sslmode=disable"),
			MaxConns:          getEnvInt32("CATALOG_DB_MAX_CONNS", 10),
			MinConns:          getEnvInt32("CATALOG_DB_MIN_CONNS", 2),
			MaxConnLifetime:   getEnvDuration("CATALOG_DB_MAX_CONN_LIFETIME", time.Hour),
			MaxConnIdleTime:   getEnvDuration("CATALOG_DB_MAX_CONN_IDLE_TIME", 30*time.Minute),
			HealthCheckPeriod: getEnvDuration("CATALOG_DB_HEALTH_CHECK_PERIOD", time.Minute),
			ConnectTimeout:    getEnvDuration("CATALOG_DB_CONNECT_TIMEOUT", 5*time.Second),
		},
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

func getEnvInt32(key string, def int32) int32 {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		log.Fatalf("invalid int value for %s: %q: %v", key, v, err)
	}
	return int32(n)
}

// getEnvDuration expects Go duration format, e.g. "5s", "30m", "1h".
func getEnvDuration(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		log.Fatalf("invalid duration value for %s: %q: %v", key, v, err)
	}
	return d
}
