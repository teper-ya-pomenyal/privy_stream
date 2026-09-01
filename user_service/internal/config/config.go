package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	UserDB         *UserDBConfig
	UserCacheDB    *UserDBCacheConfig
	TTLRefresh     time.Duration
	PrivateKeyPath string
}

func LoadConfig() *Config {

	ttlRefresh := getIntEnvOrDefault("TTL_REFRESH_TIME", 604800)

	privateKey := mustGetEnv("PRIVATE_KEY_PATH")
	return &Config{
		UserDB:         NewUserDBConfig(),
		UserCacheDB:    NewUserDBCacheConfig(),
		TTLRefresh:     time.Duration(ttlRefresh) * time.Second,
		PrivateKeyPath: privateKey,
	}
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("отсутствует необходимая переменная окружения: %s", key)
	}
	return v
}

func mustGetIntEnv(key string) int {
	strParam := mustGetEnv(key)
	intParam, err := strconv.Atoi(strParam)
	if err != nil {
		log.Fatalf("ошибка преобразования переменной %s в число: %s", key, err)
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
		log.Fatalf("ошибка преобразования переменной %s в число: %s", key, err)
	}
	return intParam
}
