package config

import "time"

type UserDBCacheConfig struct {
	DBAddress    string
	DBPassword   string
	DBNumber     int
	PoolSize     int
	MinIdleConns int
	MaxRetries   int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewUserDBCacheConfig() *UserDBCacheConfig {
	address := mustGetEnv("USER_CACHE_ADDRESS")
	password := mustGetEnv("USER_CACHE_PASSWORD")
	dbNumber := getIntEnvOrDefault("USER_CACHE_NUMBER", 0)
	poolSize := getIntEnvOrDefault("USER_CACHE_POOL_SIZE", 5)
	minIdleConns := getIntEnvOrDefault("USER_CACHE_MIN_IDLE_CONNS", 2)
	maxRetries := getIntEnvOrDefault("USER_CACHE_MAX_RETRIES", 3)
	dialTimeout := getIntEnvOrDefault("USER_CACHE_DIAL_TIMEOUT", 5)
	readTimeout := getIntEnvOrDefault("USER_CACHE_READ_TIMEOUT", 3)
	writeTimeout := getIntEnvOrDefault("USER_CACHE_WRITE_TIMEOUT", 3)

	return &UserDBCacheConfig{
		DBAddress:    address,
		DBPassword:   password,
		DBNumber:     dbNumber,
		PoolSize:     poolSize,
		MinIdleConns: minIdleConns,
		MaxRetries:   maxRetries,
		DialTimeout:  time.Duration(dialTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}
}
