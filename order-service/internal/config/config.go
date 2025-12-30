package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBURL   string
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		AppPort: getEnv("APP_PORT", "3000"),
		DBURL:   getEnv("DB_URL", ""),
	}

	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
