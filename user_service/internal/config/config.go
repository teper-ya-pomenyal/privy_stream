package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port           string
	UserDB         *UserDBConfig
	UserCacheDB    *UserDBCacheConfig
	TTLRefresh     time.Duration
	TTLAccess      time.Duration
	PrivateKeyPath string
}

func LoadConfig() *Config {

	ttlRefresh := getIntEnvOrDefault("TTL_REFRESH_TIME", 604800)
	ttlAccess := getIntEnvOrDefault("TTL_ACCESS_TIME", 900)

	port := mustGetEnv("PORT")
	privateKey := mustGetEnv("PRIVATE_KEY_PATH")
	return &Config{
		Port:           port,
		UserDB:         NewUserDBConfig(),
		UserCacheDB:    NewUserDBCacheConfig(),
		TTLRefresh:     time.Duration(ttlRefresh) * time.Second,
		TTLAccess:      time.Duration(ttlAccess) * time.Second,
		PrivateKeyPath: privateKey,
	}
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}
	return v
}

func mustGetIntEnv(key string) int {
	strParam := mustGetEnv(key)
	intParam, err := strconv.Atoi(strParam)
	if err != nil {
		log.Fatalf("failed to convert variable %s to a number: %s", key, err)
	}
	return intParam
}

func getIntEnvOrDefault(key string, defaultVal int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	intParam, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("failed to convert variable %s to a number: %s", key, err)
	}
	return intParam
}
