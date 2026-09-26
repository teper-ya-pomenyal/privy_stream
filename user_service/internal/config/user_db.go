package config

import "time"

type UserDBConfig struct {
	DSN               string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
	ConnectTimeout    time.Duration
}

func NewUserDBConfig() *UserDBConfig {
	return &UserDBConfig{
		DSN:               getEnvOrDefault("DB_DSN", "postgres://postgres:postgres@localhost:5432/users_db?sslmode=disable"),
		MaxConns:          int32(getIntEnvOrDefault("DB_MAX_CONNS", 10)),
		MinConns:          int32(getIntEnvOrDefault("DB_MIN_CONNS", 2)),
		MaxConnLifetime:   getDurationEnvOrDefault("DB_MAX_CONN_LIFETIME", time.Hour),
		MaxConnIdleTime:   getDurationEnvOrDefault("DB_MAX_CONN_IDLE_TIME", 30*time.Minute),
		HealthCheckPeriod: getDurationEnvOrDefault("DB_HEALTH_CHECK_PERIOD", time.Minute),
		ConnectTimeout:    getDurationEnvOrDefault("DB_CONNECT_TIMEOUT", 5*time.Second),
	}
}
