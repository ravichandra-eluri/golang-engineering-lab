package config

import (
	"log/slog"
	"os"
)

type Config struct {
	Port      string
	DBConn    string
	JWTSecret string
}

func Load() *Config {
	cfg := &Config{
		Port:      getEnv("PORT", "8080"),
		DBConn:    getEnv("DATABASE_URL", "postgres://localhost/dev?sslmode=disable"),
		JWTSecret: getEnv("JWT_SECRET", "change-me-in-production"),
	}
	slog.Info("config loaded", "port", cfg.Port)
	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
