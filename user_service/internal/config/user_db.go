package config

type UserDBConfig struct {
	DSN string
}

func NewUserDBConfig() *UserDBConfig {
	dsn := mustGetEnv("DB_DSN")
	return &UserDBConfig{
		DSN: dsn,
	}
}
