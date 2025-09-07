package config

import (
	"log"
	"os"
)

type Config struct {
	PostgresURL string
	RedisURL    string
	JWTSecret   string
}

func Load() *Config {
	return &Config{
		PostgresURL: getEnv("POSTGRES_URL", "postgres://user:pass@localhost:5432/otp_auth?sslmode=disable"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379/0"),
		JWTSecret:   getEnv("JWT_SECRET", "change-me"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	log.Printf("[WARN] env %s not set, using fallback", key)
	return fallback
}
