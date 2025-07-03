package config

import (
	"os"
)

type Config struct {
	DBUrl      string
	ServerAddr string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		DBUrl:      getEnv("DATABASE_URL", ""),
		ServerAddr: getEnv("SERVER_ADDR", ":8080"),
		JWTSecret:  getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
