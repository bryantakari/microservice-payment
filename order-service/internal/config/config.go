package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBPORT     string
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
	DBSSLMODE  string
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		AppPort:    getEnv("APP_PORT", "3000"),
		DBPORT:     getEnv("DB_PORT", ""),
		DBHOST:     getEnv("DB_HOST", ""),
		DBNAME:     getEnv("DB_NAME", ""),
		DBPASSWORD: getEnv("DB_PASSWORD", ""),
		DBUSER:     getEnv("DB_USER", ""),
		DBSSLMODE:  getEnv("DB_SSLMODE", ""),
	}
	return cfg
}

func (cfg *Config) CreateDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUSER,
		cfg.DBUSER,
		cfg.DBHOST,
		cfg.DBPORT,
		cfg.DBNAME,
		cfg.DBSSLMODE,
	)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
